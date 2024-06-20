package students

import (
	"fmt"

	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

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
	r.Db.Raw("SELECT * FROM student WHERE id = ? AND is_deleted = false;", id).Scan(&studentsE)
	r.Db.Raw("SELECT * FROM routine_plan WHERE id_student = ? AND is_deleted = false;", id).Scan(&routineE)
	return studentsE.ToStudent(routineE)
}

func (r *studentRepository) CreateStudent(student CreateStudentDTO) int {
	studentE := studentEntity{}
	studentE.FromCreateStudent(student)

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
		utils.Map(student.Routine, func(rp CreateStudentRoutinePlanDTO) routinePlanEntity {
			r := routinePlanEntity{}
			r.FromCreateStudentRoutinePlan(rp, studentE.ID, studentE.BasePrice)
			return r
		}),
		studentE.ID,
	)

	return studentE.ID
}

func (r *studentRepository) UpdateStudent(student UpdateStudentDTO) int {
	studentE := studentEntity{}
	studentE.FromUpdateStudent(student)
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
		utils.Map(student.Routine, func(rp UpdateStudentRoutinePlanDTO) routinePlanEntity {
			r := routinePlanEntity{}
			r.FromUpdateStudentRoutinePlan(rp, studentE.ID)
			fmt.Println(r)
			return r
		}),
	)

	return studentE.ID
}

func (r *studentRepository) UpdateStudent1(student UpdateStudentDTO) int {
	studentE := studentEntity{}
	studentE.FromUpdateStudent(student)
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
	return studentE.ID
}

func (r *studentRepository) GetRoutineByIDStudent1(idStudent int, removing *[]int) []RoutinePlan {
	routine := []routinePlanEntity{}
	args := []interface{}{idStudent}
	sql := "SELECT * FROM routine_plan WHERE id_student = ? AND is_deleted = false"
	if removing != nil {
		sql += " AND id NOT IN "
		sql += utils.SqlArray(len(*removing))
		for _, id := range *removing {
			args = append(args, id)
		}
	}
	sql += ";"
	r.Db.Raw(sql, args...).Scan(&routine)
	return utils.Map(routine, func(rp routinePlanEntity) RoutinePlan {
		return rp.ToRoutinePlan()
	})
}

func (r *studentRepository) CreateRoutinePlanByIDStudent1(idStudent int, routine []CreateStudentRoutinePlanDTO) {
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

func (r *studentRepository) DeleteRoutinePlanByIDStudent1(idStudent int, routine []int) {
	sql := fmt.Sprintf(`
    UPDATE routine_plan
    SET is_deleted = true
    WHERE id_student = ? AND id IN %s;`,
		utils.SqlArray(len(routine)),
	)
	args := []interface{}{idStudent}
	for _, id := range routine {
		args = append(args, id)
	}
	r.Db.Exec(sql, args...)
}

func (r *studentRepository) CreateRoutineByIDStudent1(routine []routinePlanEntity, idStudent int) {
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

func (r *studentRepository) CreateRoutineByIDStudent(routine []routinePlanEntity, idStudent int) {
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
			utils.SqlValues(len(orderedValues)/5, 5),
		)
		r.Db.Exec(sql, orderedValues...)
	}
}
