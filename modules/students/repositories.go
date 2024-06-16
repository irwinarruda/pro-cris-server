package students

import (
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
	r.Db.Raw("SELECT * FROM student WHERE student.is_deleted = false;").Scan(&studentsArr)
	for _, studentE := range studentsArr {
		routineE := []routinePlanEntity{}
		r.Db.Raw("SELECT * FROM routine_plan WHERE id_student = ?;", studentE.ID).Scan(&routineE)
		student := studentE.ToStudent(routineE)
		students = append(students, student)
	}
	return students
}

func (r *studentRepository) GetStudentById(id int) Student {
	studentsE := studentEntity{}
	students := Student{}
	r.Db.Raw("SELECT * FROM student WHERE student.id = ?;", id).Scan(&studentsE)
	return students
}

func (r *studentRepository) CreateStudent(student Student) int {
	studentE := student.ToStudentEntity()
	sql := `
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
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    RETURNING id;`

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

	sql = `
    INSERT INTO routine_plan(
      id_student,
      week_day,
      start_hour,
      duration,
      price
    ) VALUES`

	orderedValues := []interface{}{}
	for index, routinePlan := range student.Routine {
		if index > 0 {
			sql = sql + ","
		}
		sql = sql + "\n(?, ?, ?, ?, ?)"
		routinePlanE := routinePlan.ToRoutinePlanEntity(studentE.ID)
		orderedValues = append(
			orderedValues,
			studentE.ID,
			routinePlanE.WeekDay,
			routinePlanE.StartHour,
			routinePlanE.Duration,
			routinePlanE.Price,
		)
	}
	sql += ";"
	r.Db.Exec(sql, orderedValues...)

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
      base_price = ?
    WHERE id = ?;`
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
		studentE.ID,
	)

	sql = `
    INSERT INTO routine_plan(
      id_student,
      week_day,
      start_hour,
      duration,
      price
    ) VALUES
  `
	routineExists := []interface{}{studentE.ID}
	orderedValues := []interface{}{}
	for _, routinePlan := range student.Routine {
		routinePlanE := routinePlan.ToRoutinePlanEntity(studentE.ID)
		if routinePlanE.ID != nil {
			routineExists = append(routineExists, *routinePlanE.ID)
			continue
		}
		if len(orderedValues) > 0 {
			sql = sql + ","
		}
		sql = sql + "\n(?, ?, ?, ?, ?)"
		orderedValues = append(
			orderedValues,
			studentE.ID,
			routinePlanE.WeekDay,
			routinePlanE.StartHour,
			routinePlanE.Duration,
			routinePlanE.Price,
		)
	}
	if len(orderedValues) > 0 {
		sql += ";"
		r.Db.Exec(sql, orderedValues...)
	}

	sql = `
    SELECT id FROM routine_plan
    WHERE id_student = ?
    AND id NOT IN (`
	for index := range routineExists {
		if index > 0 {
			sql = sql + ", "
		}
		sql = sql + "?"
	}
	sql += ");"
	deletedRoutines := []int{}
	r.Db.Raw(sql, routineExists...).Scan(&deletedRoutines)

	sql = `
    UPDATE routine_plan
    SET is_deleted = true
    WHERE id IN (`

	genericDeletedRoutines := []interface{}{}
	for index := range deletedRoutines {
		genericDeletedRoutines = append(genericDeletedRoutines, deletedRoutines[index])
		if index > 0 {
			sql = sql + ", "
		}
		sql = sql + "?"
	}
	if len(deletedRoutines) > 0 {
		sql += ");"
		r.Db.Exec(sql, genericDeletedRoutines...)
	}

	return studentE.ID
}
