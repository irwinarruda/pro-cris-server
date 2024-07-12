package students

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type IStudentRepository interface {
	GetAllStudents(data GetAllStudentsDTO) []Student
	GetStudentByID(data GetStudentDTO) (Student, error)
	CreateStudent(student CreateStudentDTO) int
	UpdateStudent(student UpdateStudentDTO) (int, error)
	DeleteStudent(data DeleteStudentDTO) (int, error)
	GetRoutineID(idStudent int, excluded ...int) []int
	CreateRoutine(idStudent int, routinePlan ...CreateStudentRoutinePlanDTO)
	DeleteRoutine(idStudent int, routine ...int)
	ResetStudents()
}

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
	PaymentStyle             PaymentStyle
	PaymentType              PaymentType
	PaymentTypeValue         *float64
	SettlementStyle          SettlementStyle
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

func (s *DbStudent) FromCreateStudent(student CreateStudentDTO) {
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

func (s *DbStudent) FromUpdateStudent(student UpdateStudentDTO) {
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

func (s *DbStudent) ToStudent(dbRoutine []DbRoutinePlan) Student {
	var coordinate *models.Coordinate
	if s.HouseCoordinateLatitude != nil && s.HouseCoordinateLongitude != nil {
		coordinate = &models.Coordinate{
			Latitude:  *s.HouseCoordinateLatitude,
			Longitude: *s.HouseCoordinateLongitude,
		}
	}
	return Student{
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
		Routine: utils.Map(dbRoutine, func(rp DbRoutinePlan, _ int) RoutinePlan {
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

func (r *DbRoutinePlan) ToRoutinePlan() RoutinePlan {
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
