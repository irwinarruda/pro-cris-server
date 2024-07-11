package students

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/models"
)

type PaymentStyle = string

const (
	Upfront PaymentStyle = "Upfront"
	Later   PaymentStyle = "Later"
)

func GetPaymentStyles() []PaymentStyle { return []PaymentStyle{Upfront, Later} }

type PaymentType = string

const (
	Fixed    PaymentType = "Fixed"
	Variable PaymentType = "Variable"
)

func GetPaymentTypes() []PaymentType { return []PaymentType{Fixed, Variable} }

type SettlementStyle = string

const (
	Appointments SettlementStyle = "Appointments"
	Weekly       SettlementStyle = "Weekly"
	Monthly      SettlementStyle = "Monthly"
)

func GetSettlementStyles() []SettlementStyle { return []SettlementStyle{Appointments, Weekly, Monthly} }

type Student struct {
	ID                   int                `json:"id"`
	Name                 string             `json:"name"`
	BirthDay             *string            `json:"birthDay"`
	DisplayColor         string             `json:"displayColor"`
	Picture              *string            `json:"picture"`
	Gender               *models.Gender     `json:"gender"`
	ParentName           *string            `json:"parentName"`
	ParentPhoneNumber    *string            `json:"parentPhoneNumber"`
	PaymentStyle         PaymentStyle       `json:"paymentStyle"`
	PaymentType          PaymentType        `json:"paymentType"`
	PaymentTypeValue     *float64           `json:"paymentTypeValue"`
	SettlementStyle      SettlementStyle    `json:"settlementStyle"`
	SettlementStyleValue *int               `json:"settlementStyleValue"`
	SettlementStyleDay   *int               `json:"settlementStyleDay"`
	HouseAddress         *string            `json:"houseAddress"`
	HouseIdentifier      *string            `json:"hoseIdentifier"`
	HouseCoordinate      *models.Coordinate `json:"houseCoordinate"`
	Routine              []RoutinePlan      `json:"routine"`
	IsDeleted            bool               `json:"isDeleted"`
	CreatedAt            time.Time          `json:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt"`
}

func (s *Student) ToStudentEntity() StudentEntity {
	var latitude *float64
	var longitude *float64
	if s.HouseCoordinate != nil {
		latitude = &s.HouseCoordinate.Latitude
		longitude = &s.HouseCoordinate.Longitude
	}
	return StudentEntity{
		ID:                       s.ID,
		Name:                     s.Name,
		BirthDay:                 s.BirthDay,
		DisplayColor:             s.DisplayColor,
		Picture:                  s.Picture,
		Gender:                   s.Gender,
		ParentName:               s.ParentName,
		ParentPhoneNumber:        s.ParentPhoneNumber,
		PaymentStyle:             s.PaymentStyle,
		PaymentType:              s.PaymentType,
		PaymentTypeValue:         s.PaymentTypeValue,
		SettlementStyle:          s.SettlementStyle,
		SettlementStyleValue:     s.SettlementStyleValue,
		SettlementStyleDay:       s.SettlementStyleDay,
		HouseAddress:             s.HouseAddress,
		HouseIdentifier:          s.HouseIdentifier,
		HouseCoordinateLatitude:  latitude,
		HouseCoordinateLongitude: longitude,
		IsDeleted:                s.IsDeleted,
		CreatedAt:                s.CreatedAt,
	}
}

type RoutinePlan struct {
	ID        int            `json:"id"`
	WeekDay   models.WeekDay `json:"weekDay"`
	StartHour int            `json:"startHour"` // milisseconds
	Duration  int            `json:"duration"`  // milisseconds
	Price     float64        `json:"price"`
	IsDeleted bool           `json:"isDeleted"`
	CreatedAt time.Time      `json:"createdAt"`
}

func (r *RoutinePlan) ToRoutinePlanEntity(idStudent int) routinePlanEntity {
	return routinePlanEntity{
		ID:        &r.ID,
		IDStudent: idStudent,
		WeekDay:   r.WeekDay,
		StartHour: r.StartHour,
		Duration:  r.Duration,
		Price:     r.Price,
		IsDeleted: r.IsDeleted,
		CreatedAt: r.CreatedAt,
	}
}
