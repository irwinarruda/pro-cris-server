package students

import (
	"fmt"
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type studentEntity struct {
	ID                       int
	Name                     string
	BirthDay                 *string
	DisplayColor             string
	Picture                  *string
	ParentName               *string
	ParentPhoneNumber        *string
	HouseAddress             *string
	HouseIdentifier          *string
	HouseCoordinateLatitude  *float64
	HouseCoordinateLongitude *float64
	BasePrice                float64
	IsDeleted                bool
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

func (s *studentEntity) ToStudent(routineEntity []routinePlanEntity) Student {
	var coordinate *models.Coordinate
	if s.HouseCoordinateLatitude != nil && s.HouseCoordinateLongitude != nil {
		coordinate = &models.Coordinate{
			Latitude:  *s.HouseCoordinateLatitude,
			Longitude: *s.HouseCoordinateLongitude,
		}
	}
	return Student{
		ID:                s.ID,
		Name:              s.Name,
		BirthDay:          s.BirthDay,
		DisplayColor:      s.DisplayColor,
		Picture:           s.Picture,
		ParentName:        s.ParentName,
		ParentPhoneNumber: s.ParentPhoneNumber,
		HouseAddress:      s.HouseAddress,
		HouseIdentifier:   s.HouseIdentifier,
		HouseCoordinate:   coordinate,
		BasePrice:         s.BasePrice,
		IsDeleted:         s.IsDeleted,
		CreatedAt:         s.CreatedAt,
		UpdatedAt:         s.UpdatedAt,
		Routine: utils.Map(routineEntity, func(rp routinePlanEntity) RoutinePlan {
			return rp.ToRoutinePlan()
		}),
	}
}

type routinePlanEntity struct {
	ID        *int
	IdStudent int
	WeekDay   models.WeekDay
	StartHour int
	Duration  int
	Price     float64
	IsDeleted bool
	CreatedAt time.Time
}

func (r *routinePlanEntity) ToRoutinePlan() RoutinePlan {
	var id *int
	if r.ID != nil {
		id = r.ID
	}
	return RoutinePlan{
		ID:        *id,
		WeekDay:   r.WeekDay,
		StartHour: r.StartHour,
		Duration:  r.Duration,
		Price:     r.Price,
		IsDeleted: r.IsDeleted,
		CreatedAt: r.CreatedAt,
	}
}

type studentRepository struct {
	Db configs.Db `inject:"db"`
}

func newStudentRepository() *studentRepository {
	return configs.ResolveInject(&studentRepository{})
}

func (r *studentRepository) GetAllStudents() []Student {
	studentsArr := []studentEntity{}
	students := []Student{}
	r.Db.Raw("SELECT * FROM student WHERE is_deleted = false;").Scan(&studentsArr)
	for _, studentE := range studentsArr {
		routineE := []routinePlanEntity{}
		r.Db.Raw("SELECT * FROM routine_plan WHERE id_student = ? AND is_deleted = false;", studentE.ID).Scan(&routineE)
		student := studentE.ToStudent(routineE)
		students = append(students, student)
	}
	return students
}

func (r *studentRepository) GetStudentByID(id int) Student {
	studentsE := studentEntity{}
	routineE := []routinePlanEntity{}
	r.Db.Raw("SELECT * FROM student WHERE id = ?;", id).Scan(&studentsE)
	r.Db.Raw("SELECT * FROM routine_plan WHERE id_student = ? AND is_deleted = false;", id).Scan(&routineE)
	return studentsE.ToStudent(routineE)
}

func (r *studentRepository) CreateStudent(student Student) int {
	studentE := student.ToStudentEntity()
	sql := fmt.Sprintf(`
    INSERT INTO student(
      name,
      birth_day,
      display_color,
      picture,
      parent_name,
      parent_phone_number,
      house_address,
      house_identifier,
      house_coordinate_latitude,
      house_coordinate_longitude,
      base_price
    )
    %s
    RETURNING id;`,
		utils.SqlValues(1, 11),
	)

	r.Db.Raw(
		sql,
		studentE.Name,
		studentE.BirthDay,
		studentE.DisplayColor,
		studentE.Picture,
		studentE.ParentName,
		studentE.ParentPhoneNumber,
		studentE.HouseAddress,
		studentE.HouseIdentifier,
		studentE.HouseCoordinateLatitude,
		studentE.HouseCoordinateLongitude,
		studentE.BasePrice,
	).Scan(&studentE.ID)

	r.CreateRoutineByIDStudent(
		studentE.ID,
		utils.Map(student.Routine, func(rp RoutinePlan) routinePlanEntity {
			return rp.ToRoutinePlanEntity(studentE.ID)
		}),
	)

	return studentE.ID
}

func (r *studentRepository) UpdateStudent(student UpdateStudentDTO) int {
	studentE := student.ToStudentEntity()
	sql := `
    UPDATE student
    SET
      name = ?,
      birth_day = ?,
      display_color = ?,
      picture = ?,
      parent_name = ?,
      parent_phone_number = ?,
      house_address = ?,
      house_identifier = ?,
      house_coordinate_latitude = ?,
      house_coordinate_longitude = ?,
      base_price = ?,
      updated_at = now()
    WHERE id = ?;`
	r.Db.Exec(
		sql,
		studentE.Name,
		studentE.BirthDay,
		studentE.DisplayColor,
		studentE.Picture,
		studentE.ParentName,
		studentE.ParentPhoneNumber,
		studentE.HouseAddress,
		studentE.HouseIdentifier,
		studentE.HouseCoordinateLatitude,
		studentE.HouseCoordinateLongitude,
		studentE.BasePrice,
		studentE.ID,
	)

	r.UpdateRoutineByIDStudent(
		studentE.ID,
		utils.Map(student.Routine, func(rp UpdateStudentRoutineDTO) routinePlanEntity {
			return rp.ToRoutinePlanEntity(studentE.ID)
		}),
	)

	return studentE.ID
}

func (r *studentRepository) DeleteStudentByID(id int) {
	sql := `
    UPDATE student
    SET is_deleted = true
    WHERE id = ?;`
	r.Db.Exec(sql, id)
	sql = `
    UPDATE routine_plan
    SET is_deleted = true
    WHERE id_student = ?;`
	r.Db.Exec(sql, id)
}

func (r *studentRepository) CreateRoutineByIDStudent(idStudent int, routine []routinePlanEntity) {
	orderedValues := []interface{}{}
	sql := fmt.Sprintf(`
    INSERT INTO routine_plan(
      id_student,
      week_day,
      start_hour,
      duration,
      price
    ) %s;`,
		utils.SqlValues(len(routine), 5, func(index int) {
			routinePlanE := routine[index]
			orderedValues = append(
				orderedValues,
				idStudent,
				routinePlanE.WeekDay,
				routinePlanE.StartHour,
				routinePlanE.Duration,
				routinePlanE.Price,
			)
		}),
	)
	r.Db.Exec(sql, orderedValues...)
}

func (r *studentRepository) UpdateRoutineByIDStudent(idStudent int, routine []routinePlanEntity) {
	orderedValues := []interface{}{}
	routineExists := []interface{}{idStudent}
	for index := range routine {
		routinePlanE := routine[index]
		if routinePlanE.ID == nil {
			orderedValues = append(
				orderedValues,
				idStudent,
				routinePlanE.WeekDay,
				routinePlanE.StartHour,
				routinePlanE.Duration,
				routinePlanE.Price,
			)
			continue
		}
		routineExists = append(routineExists, *routinePlanE.ID)
	}

	sql := `
      SELECT id FROM routine_plan
      WHERE id_student = ?`
	if len(routineExists) > 1 {
		sql += " AND id NOT IN "
		sql += utils.SqlArray(len(routineExists) - 1)
	}
	sql += ";"
	deletedRoutines := []interface{}{}
	r.Db.Raw(sql, routineExists...).Scan(&deletedRoutines)

	if len(deletedRoutines) > 0 {
		genericDeletedRoutines := []interface{}{}
		sql = fmt.Sprintf(`
      UPDATE routine_plan
      SET is_deleted = true
      WHERE id IN %s;`,
			utils.SqlArray(len(deletedRoutines), func(index int) {
				genericDeletedRoutines = append(genericDeletedRoutines, deletedRoutines[index])
			}),
		)
		sql += ";"
		r.Db.Exec(sql, genericDeletedRoutines...)
	}

	if len(orderedValues) > 0 {
		sql = fmt.Sprintf(`
      INSERT INTO routine_plan(
        id_student,
        week_day,
        start_hour,
        duration,
        price
      ) %s;`,
			utils.SqlValues(len(orderedValues), 5),
		)
		r.Db.Exec(sql, orderedValues...)
	}
}
