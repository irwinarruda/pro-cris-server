package students

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type PaymentStyle string

const (
	PaymentStyleUpfront PaymentStyle = "Upfront"
	PaymentStyleLater   PaymentStyle = "Later"
)

func (s *PaymentStyle) UnmarshalJSON(b []byte) (err error) {
	return utils.UnmarshalEnum(s, GetPaymentStylesString(), b)
}

func (p PaymentStyle) String() string { return string(p) }

func GetPaymentStylesString() []string {
	return []string{PaymentStyleUpfront.String(), PaymentStyleLater.String()}
}

type PaymentType string

const (
	PaymentTypeFixed    PaymentType = "Fixed"
	PaymentTypeVariable PaymentType = "Variable"
)

func (p *PaymentType) UnmarshalJSON(b []byte) (err error) {
	return utils.UnmarshalEnum(p, GetPaymentTypesString(), b)
}

func (p PaymentType) String() string { return string(p) }

func GetPaymentTypesString() []string {
	return []string{PaymentTypeFixed.String(), PaymentTypeVariable.String()}
}

type SettlementStyle string

const (
	SettlementStyleAppointments SettlementStyle = "Appointments"
	SettlementStyleWeekly       SettlementStyle = "Weekly"
	SettlementStyleMonthly      SettlementStyle = "Monthly"
)

func (s *SettlementStyle) UnmarshalJSON(b []byte) (err error) {
	return utils.UnmarshalEnum(s, GetSettlementStylesString(), b)
}

func (p SettlementStyle) String() string { return string(p) }

func GetSettlementStylesString() []string {
	return []string{SettlementStyleAppointments.String(), SettlementStyleWeekly.String(), SettlementStyleMonthly.String()}
}

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

type RoutinePlan struct {
	ID        int            `json:"id"`
	WeekDay   models.WeekDay `json:"weekDay"`
	StartHour int            `json:"startHour"` // milisseconds
	Duration  int            `json:"duration"`  // milisseconds
	Price     float64        `json:"price"`
	IsDeleted bool           `json:"isDeleted"`
	CreatedAt time.Time      `json:"createdAt"`
}
