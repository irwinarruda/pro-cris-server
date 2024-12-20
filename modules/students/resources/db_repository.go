package studentsresources

import (
	"fmt"
	"net/http"
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type DbStudent struct {
	ID                       int
	IDAccount                int
	Name                     string
	BirthDay                 *string
	DisplayColor             string
	Picture                  *string
	Gender                   *string
	ParentName               *string
	ParentPhoneNumber        *string
	PaymentStyle             models.PaymentStyle
	PaymentType              models.PaymentType
	PaymentTypeValue         *float64
	SettlementStyle          models.SettlementStyle
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
	s.IDAccount = student.IDAccount
	s.Name = student.Name
	s.BirthDay = student.BirthDay
	s.DisplayColor = student.DisplayColor
	s.Picture = student.Picture
	gender := student.Gender.String()
	s.Gender = &gender
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
	s.IDAccount = student.IDAccount
	s.Name = student.Name
	s.BirthDay = student.BirthDay
	s.DisplayColor = student.DisplayColor
	s.Picture = student.Picture
	gender := student.Gender.String()
	s.Gender = &gender
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
		ID:                s.ID,
		Name:              s.Name,
		BirthDay:          s.BirthDay,
		DisplayColor:      s.DisplayColor,
		Picture:           s.Picture,
		Gender:            (*models.Gender)(s.Gender),
		ParentName:        s.ParentName,
		ParentPhoneNumber: s.ParentPhoneNumber,
		HouseAddress:      s.HouseAddress,
		HouseIdentifier:   s.HouseIdentifier,
		HouseCoordinate:   coordinate,
		IsDeleted:         s.IsDeleted,
		CreatedAt:         s.CreatedAt,
		UpdatedAt:         s.UpdatedAt,
		Routine: utils.Map(dbRoutine, func(rp DbRoutinePlan, _ int) students.RoutinePlan {
			return rp.ToRoutinePlan()
		}),
		SettlementOptions: models.SettlementOptions{
			PaymentStyle:         s.PaymentStyle,
			PaymentType:          s.PaymentType,
			PaymentTypeValue:     s.PaymentTypeValue,
			SettlementStyle:      s.SettlementStyle,
			SettlementStyleValue: s.SettlementStyleValue,
			SettlementStyleDay:   s.SettlementStyleDay,
		},
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

func (r *DbStudentRepository) GetAllStudents(data students.GetAllStudentsDTO) ([]students.Student, error) {
	studentsArr := []DbStudent{}
	students := []students.Student{}
	result := r.Db.Raw(`SELECT * FROM "student" WHERE id_account = ? AND is_deleted = false;`, data.IDAccount).Scan(&studentsArr)
	if result.Error != nil {
		return students, result.Error
	}
	for _, studentE := range studentsArr {
		routineE := []DbRoutinePlan{}
		result = r.Db.Raw(`SELECT * FROM "routine_plan" WHERE id_student = ? AND is_deleted = false;`, studentE.ID).Scan(&routineE)
		if result.Error != nil {
			return students, result.Error
		}
		student := studentE.ToStudent(routineE)
		students = append(students, student)
	}
	return students, nil
}

func (r *DbStudentRepository) GetStudentByID(data students.GetStudentDTO) (students.Student, error) {
	studentsE := []DbStudent{}
	result := r.Db.Raw(`SELECT * FROM "student" WHERE id_account = ? AND id = ? AND is_deleted = false;`, data.IDAccount, data.ID).Scan(&studentsE)
	if result.Error != nil {
		return students.Student{}, result.Error
	}
	if len(studentsE) == 0 {
		return students.Student{}, utils.NewAppError("Student not found.", true, http.StatusBadRequest)
	}
	routineE := []DbRoutinePlan{}
	result = r.Db.Raw(`SELECT * FROM "routine_plan" WHERE id_student = ? AND is_deleted = false;`, data.ID).Scan(&routineE)
	if result.Error != nil {
		return students.Student{}, result.Error
	}
	return studentsE[0].ToStudent(routineE), nil
}

func (r *DbStudentRepository) CreateStudent(student students.CreateStudentDTO) int {
	studentE := DbStudent{}
	studentE.FromCreateStudent(student)
	sql := fmt.Sprintf(`
    INSERT INTO "student"(
      id_account,
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
	result := r.Db.Raw(
		sql,
		studentE.IDAccount,
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
	if result.Error != nil {
		return 0
	}

	r.CreateRoutine(studentE.ID, student.Routine...)
	return studentE.ID
}

func (r *DbStudentRepository) UpdateStudent(student students.UpdateStudentDTO) (int, error) {
	studentE := DbStudent{}
	studentE.FromUpdateStudent(student)
	sql := `
    UPDATE "student"
    SET
      id_account = ?,
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
      updated_at = CURRENT_TIMESTAMP
    WHERE id = ?`
	result := r.Db.Exec(
		sql,
		studentE.IDAccount,
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
	)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, utils.NewAppError("Student not found.", true, http.StatusBadRequest)
	}
	return studentE.ID, nil
}

func (r *DbStudentRepository) DeleteStudent(data students.DeleteStudentDTO) (int, error) {
	sql := `
    UPDATE "student"
    SET
      is_deleted = true,
      updated_at = CURRENT_TIMESTAMP
    WHERE id_account = ?
    AND id = ?
    RETURNING id;`
	result := r.Db.Exec(sql, data.IDAccount, data.ID)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, utils.NewAppError("Student not found.", true, http.StatusBadRequest)
	}
	sql = `
    UPDATE "routine_plan"
    SET is_deleted = true
    WHERE id_student = ?;`
	result = r.Db.Exec(sql, data.ID)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, utils.NewAppError("Student not found.", true, http.StatusBadRequest)
	}
	return data.ID, nil
}

// Get Routine from a student.
//
// 'excluded' is a list of ids that should be excluded ([]int).
func (r *DbStudentRepository) GetRoutineID(idStudent int, excluded ...int) []int {
	routine := []int{}
	args := []interface{}{idStudent}
	sql := `SELECT id FROM "routine_plan" WHERE id_student = ? AND is_deleted = false`
	if len(excluded) > 0 {
		sql += ` AND id NOT IN `
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
    INSERT INTO "routine_plan"(
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
func (r *DbStudentRepository) DeleteAllRoutine(idStudent int, except ...int) {
	sql := `
    UPDATE "routine_plan"
    SET is_deleted = true
    WHERE id_student = ?`
	args := []interface{}{idStudent}
	if len(except) > 0 {
		sql += ` AND id NOT IN `
		sql += utils.SqlArray(len(except))
		for _, id := range except {
			args = append(args, id)
		}
	}
	sql += ";"
	r.Db.Exec(sql, args...)
}

func (r *DbStudentRepository) ResetStudents() {
	r.Db.Exec(`DELETE FROM "student";`)
	r.Db.Exec(`ALTER SEQUENCE student_id_seq RESTART WITH 1;`)
	r.Db.Exec(`DELETE FROM "routine_plan";`)
	r.Db.Exec(`ALTER SEQUENCE routine_plan_id_seq RESTART WITH 1;`)
}
