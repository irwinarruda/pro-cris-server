package students

import (
	"fmt"

	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type StudentRepository struct {
	Db configs.Db `inject:"db"`
}

func (r *StudentRepository) GetAllStudents() []Student {
	studentsArr := []StudentEntity{}
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

func (r *StudentRepository) GetStudentByID(id int) (Student, error) {
	studentsE := []StudentEntity{}
	r.Db.Raw("SELECT * FROM student WHERE id = ? AND is_deleted = false;", id).Scan(&studentsE)
	if len(studentsE) == 0 {
		return Student{}, utils.NewAppError("Student not found.", true, nil)
	}
	routineE := []routinePlanEntity{}
	r.Db.Raw("SELECT * FROM routine_plan WHERE id_student = ? AND is_deleted = false;", id).Scan(&routineE)
	return studentsE[0].ToStudent(routineE), nil
}

func (r *StudentRepository) CreateStudent(student CreateStudentDTO) int {
	studentE := StudentEntity{}
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
	r.CreateRoutine(studentE.ID, student.Routine...)
	return studentE.ID
}

func (r *StudentRepository) UpdateStudent(student UpdateStudentDTO) (int, error) {
	var id *int
	studentE := StudentEntity{}
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
    WHERE id = ?
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
		studentE.ID,
	).Scan(&id)
	if id == nil {
		return 0, utils.NewAppError("Student not found.", true, nil)
	}
	return *id, nil
}

func (r *StudentRepository) DeleteStudent(id int) (int, error) {
	var idStudent *int
	sql := `
    UPDATE student
    SET is_deleted = true
    WHERE id = ?
    RETURNING id;`
	r.Db.Raw(sql, id).Scan(&idStudent)
	if idStudent == nil {
		return 0, utils.NewAppError("Student not found.", true, nil)
	}
	sql = `
    UPDATE routine_plan
    SET is_deleted = true
    WHERE id_student = ?;`
	r.Db.Exec(sql, id)
	return id, nil
}

// Get Routine from a student.
//
// 'excluded' is a list of ids that should be excluded ([]int).
func (r *StudentRepository) GetRoutineID(idStudent int, excluded ...int) []int {
	routine := []int{}
	args := []interface{}{idStudent}
	sql := "SELECT id FROM routine_plan WHERE id_student = ? AND is_deleted = false"
	if excluded != nil && len(excluded) > 0 {
		sql += " AND id NOT IN "
		sql += utils.SqlArray(len(excluded))
		for _, id := range excluded {
			args = append(args, id)
		}
	}
	sql += ";"
	r.Db.Raw(sql, args...).Scan(&routine)
	return routine
}

// Create a list of RoutinePlan from a student.
//
// 'routinePlan' can be either one or more items.
func (r *StudentRepository) CreateRoutine(idStudent int, routinePlan ...CreateStudentRoutinePlanDTO) {
	if len(routinePlan) == 0 {
		return
	}

	orderedValues := []interface{}{}
	sql := fmt.Sprintf(`
    INSERT INTO routine_plan(
      id_student,
      week_day,
      start_hour,
      duration,
      price
    ) %s;`,
		utils.SqlValues(len(routinePlan), 5, func(index int) {
			routinePlanE := routinePlan[index]
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

// Delete a list of RoutinePlan from a student.
//
// 'routine' is a list of ids that should be deleted.
func (r *StudentRepository) DeleteRoutine(idStudent int, routine ...int) {
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

func (r *StudentRepository) ResetStudents() {
	sql := "DELETE FROM student;"
	r.Db.Exec(sql)
}
