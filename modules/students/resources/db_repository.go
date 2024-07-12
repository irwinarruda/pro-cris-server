package studentsresources

import (
	"fmt"
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type DbStudent struct {
	ID                       int
	IDUser                   int
	Name                     string
	BirthDay                 *string
	DisplayColor             string
	Picture                  *string
	Gender                   *string
	ParentName               *string
	ParentPhoneNumber        *string
	PaymentStyle             students.PaymentStyle
	PaymentType              students.PaymentType
	PaymentTypeValue         *float64
	SettlementStyle          students.SettlementStyle
	SettlementStyleValue     *int
	SettlementStyleDay       *int
	HouseAddress             *string
	HouseIdentifier          *string
	HouseCoordinateLatitude  *float64
	HouseCoordinateLongitude *float64
	IsDeleted                bool
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

func (s *DbStudent) FromCreateStudent(student students.CreateStudentDTO) {
	var latitude *float64
	var longitude *float64
	if student.HouseCoordinate != nil {
		latitude = &student.HouseCoordinate.Latitude
		longitude = &student.HouseCoordinate.Longitude
	}
	s.IDUser = student.IDUser
	s.Name = student.Name
	s.BirthDay = student.BirthDay
	s.DisplayColor = student.DisplayColor
	s.Picture = student.Picture
	s.Gender = student.Gender
	s.ParentName = student.ParentName
	s.ParentPhoneNumber = student.ParentPhoneNumber
	s.PaymentStyle = student.PaymentStyle
	s.PaymentType = student.PaymentType
	s.PaymentTypeValue = student.PaymentTypeValue
	s.SettlementStyle = student.SettlementStyle
	s.SettlementStyleValue = student.SettlementStyleValue
	s.SettlementStyleDay = student.SettlementStyleDay
	s.HouseAddress = student.HouseAddress
	s.HouseIdentifier = student.HouseIdentifier
	s.HouseCoordinateLatitude = latitude
	s.HouseCoordinateLongitude = longitude
}

func (s *DbStudent) FromUpdateStudent(student students.UpdateStudentDTO) {
	var latitude *float64
	var longitude *float64
	if student.HouseCoordinate != nil {
		latitude = &student.HouseCoordinate.Latitude
		longitude = &student.HouseCoordinate.Longitude
	}
	s.ID = student.ID
	s.IDUser = student.IDUser
	s.Name = student.Name
	s.BirthDay = student.BirthDay
	s.DisplayColor = student.DisplayColor
	s.Picture = student.Picture
	s.Gender = student.Gender
	s.ParentName = student.ParentName
	s.ParentPhoneNumber = student.ParentPhoneNumber
	s.PaymentStyle = student.PaymentStyle
	s.PaymentType = student.PaymentType
	s.PaymentTypeValue = student.PaymentTypeValue
	s.SettlementStyle = student.SettlementStyle
	s.SettlementStyleValue = student.SettlementStyleValue
	s.SettlementStyleDay = student.SettlementStyleDay
	s.HouseAddress = student.HouseAddress
	s.HouseIdentifier = student.HouseIdentifier
	s.HouseCoordinateLatitude = latitude
	s.HouseCoordinateLongitude = longitude
	s.UpdatedAt = time.Now()
}

func (s *DbStudent) ToStudent(dbRoutine []DbRoutinePlan) students.Student {
	var coordinate *models.Coordinate
	if s.HouseCoordinateLatitude != nil && s.HouseCoordinateLongitude != nil {
		coordinate = &models.Coordinate{
			Latitude:  *s.HouseCoordinateLatitude,
			Longitude: *s.HouseCoordinateLongitude,
		}
	}
	return students.Student{
		ID:                   s.ID,
		Name:                 s.Name,
		BirthDay:             s.BirthDay,
		DisplayColor:         s.DisplayColor,
		Picture:              s.Picture,
		Gender:               s.Gender,
		ParentName:           s.ParentName,
		ParentPhoneNumber:    s.ParentPhoneNumber,
		PaymentStyle:         s.PaymentStyle,
		PaymentType:          s.PaymentType,
		PaymentTypeValue:     s.PaymentTypeValue,
		SettlementStyle:      s.SettlementStyle,
		SettlementStyleValue: s.SettlementStyleValue,
		SettlementStyleDay:   s.SettlementStyleDay,
		HouseAddress:         s.HouseAddress,
		HouseIdentifier:      s.HouseIdentifier,
		HouseCoordinate:      coordinate,
		IsDeleted:            s.IsDeleted,
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
		Routine: utils.Map(dbRoutine, func(rp DbRoutinePlan, _ int) students.RoutinePlan {
			return rp.ToRoutinePlan()
		}),
	}
}

type DbRoutinePlan struct {
	ID        *int
	IDStudent int
	WeekDay   models.WeekDay
	StartHour int
	Duration  int
	Price     float64
	IsDeleted bool
	CreatedAt time.Time
}

func (r *DbRoutinePlan) ToRoutinePlan() students.RoutinePlan {
	var id *int
	if r.ID != nil {
		id = r.ID
	}
	return students.RoutinePlan{
		ID:        *id,
		WeekDay:   r.WeekDay,
		StartHour: r.StartHour,
		Duration:  r.Duration,
		Price:     r.Price,
		IsDeleted: r.IsDeleted,
		CreatedAt: r.CreatedAt,
	}
}

type DbStudentRepository struct {
	Db configs.Db `inject:"db"`
}

func NewDbStudentRepository() *DbStudentRepository {
	return proinject.Resolve(&DbStudentRepository{})
}

func (r *DbStudentRepository) GetAllStudents(data students.GetAllStudentsDTO) []students.Student {
	studentsArr := []DbStudent{}
	students := []students.Student{}
	r.Db.Raw("SELECT * FROM student WHERE id_user = ? AND is_deleted = false;", data.IDUser).Scan(&studentsArr)
	for _, studentE := range studentsArr {
		routineE := []DbRoutinePlan{}
		r.Db.Raw("SELECT * FROM routine_plan WHERE id_student = ? AND is_deleted = false;", studentE.ID).Scan(&routineE)
		student := studentE.ToStudent(routineE)
		students = append(students, student)
	}
	return students
}

func (r *DbStudentRepository) GetStudentByID(data students.GetStudentDTO) (students.Student, error) {
	studentsE := []DbStudent{}
	r.Db.Raw("SELECT * FROM student WHERE id_user = ? AND id = ? AND is_deleted = false;", data.IDUser, data.ID).Scan(&studentsE)
	if len(studentsE) == 0 {
		return students.Student{}, utils.NewAppError("Student not found.", true, nil)
	}
	routineE := []DbRoutinePlan{}
	r.Db.Raw("SELECT * FROM routine_plan WHERE id_student = ? AND is_deleted = false;", data.ID).Scan(&routineE)
	return studentsE[0].ToStudent(routineE), nil
}

func (r *DbStudentRepository) CreateStudent(student students.CreateStudentDTO) int {
	studentE := DbStudent{}
	studentE.FromCreateStudent(student)
	sql := fmt.Sprintf(`
    INSERT INTO student(
      id_user,
      name,
      birth_day,
      display_color,
      picture,
      gender,
      parent_name,
      parent_phone_number,
      payment_style,
      payment_type,
      payment_type_value,
      settlement_style,
      settlement_style_value,
      settlement_style_day,
      house_address,
      house_identifier,
      house_coordinate_latitude,
      house_coordinate_longitude
    )
    %s
    RETURNING id;`,
		utils.SqlValues(1, 18),
	)
	r.Db.Raw(
		sql,
		studentE.IDUser,
		studentE.Name,
		studentE.BirthDay,
		studentE.DisplayColor,
		studentE.Picture,
		studentE.Gender,
		studentE.ParentName,
		studentE.ParentPhoneNumber,
		studentE.PaymentStyle,
		studentE.PaymentType,
		studentE.PaymentTypeValue,
		studentE.SettlementStyle,
		studentE.SettlementStyleValue,
		studentE.SettlementStyleDay,
		studentE.HouseAddress,
		studentE.HouseIdentifier,
		studentE.HouseCoordinateLatitude,
		studentE.HouseCoordinateLongitude,
	).Scan(&studentE.ID)
	r.CreateRoutine(studentE.ID, student.Routine...)
	return studentE.ID
}

func (r *DbStudentRepository) UpdateStudent(student students.UpdateStudentDTO) (int, error) {
	var id *int
	studentE := DbStudent{}
	studentE.FromUpdateStudent(student)
	sql := `
    UPDATE student
    SET
      id_user = ?,
      name = ?,
      birth_day = ?,
      display_color = ?,
      picture = ?,
      gender = ?,
      parent_name = ?,
      parent_phone_number = ?,
      payment_style = ?,
      payment_type = ?,
      payment_type_value = ?,
      settlement_style = ?,
      settlement_style_value = ?,
      settlement_style_day = ?,
      house_address = ?,
      house_identifier = ?,
      house_coordinate_latitude = ?,
      house_coordinate_longitude = ?,
      updated_at = now()
    WHERE id = ?
    RETURNING id;`
	r.Db.Raw(
		sql,
		studentE.IDUser,
		studentE.Name,
		studentE.BirthDay,
		studentE.DisplayColor,
		studentE.Picture,
		studentE.Gender,
		studentE.ParentName,
		studentE.ParentPhoneNumber,
		studentE.PaymentStyle,
		studentE.PaymentType,
		studentE.PaymentTypeValue,
		studentE.SettlementStyle,
		studentE.SettlementStyleValue,
		studentE.SettlementStyleDay,
		studentE.HouseAddress,
		studentE.HouseIdentifier,
		studentE.HouseCoordinateLatitude,
		studentE.HouseCoordinateLongitude,
		studentE.ID,
	).Scan(&id)
	if id == nil {
		return 0, utils.NewAppError("Student not found.", true, nil)
	}
	return *id, nil
}

func (r *DbStudentRepository) DeleteStudent(data students.DeleteStudentDTO) (int, error) {
	var idStudent *int
	sql := `
    UPDATE student
    SET is_deleted = true
    WHERE id_user = ?
    AND id = ?
    RETURNING id;`
	r.Db.Raw(sql, data.IDUser, data.ID).Scan(&idStudent)
	if idStudent == nil {
		return 0, utils.NewAppError("Student not found.", true, nil)
	}
	sql = `
    UPDATE routine_plan
    SET is_deleted = true
    WHERE id_student = ?;`
	r.Db.Exec(sql, data.ID)
	return data.ID, nil
}

// Get Routine from a student.
//
// 'excluded' is a list of ids that should be excluded ([]int).
func (r *DbStudentRepository) GetRoutineID(idStudent int, excluded ...int) []int {
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
func (r *DbStudentRepository) CreateRoutine(idStudent int, routinePlan ...students.CreateStudentRoutinePlanDTO) {
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
func (r *DbStudentRepository) DeleteRoutine(idStudent int, routine ...int) {
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

func (r *DbStudentRepository) ResetStudents() {
	r.Db.Exec("DELETE FROM student;")
	r.Db.Exec("ALTER SEQUENCE student_id_seq RESTART WITH 1;")
}
