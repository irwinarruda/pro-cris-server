package students

import "github.com/irwinarruda/pro-cris-server/shared/models"

type GetAllStudentsDTO struct {
	IDUser int `json:"idUser"`
}

type GetStudentDTO struct {
	IDUser int `json:"idUser"`
	ID     int `json:"id"`
}

type CreateStudentDTO struct {
	IDUser               int                           `json:"idUser"`
	Name                 string                        `json:"name" validate:"required"`
	BirthDay             *string                       `json:"birthDay" validate:"omitempty,datetime"`
	DisplayColor         string                        `json:"displayColor" validate:"omitempty,hexcolor"`
	Gender               *models.Gender                `json:"gender"`
	Picture              *string                       `json:"picture" validate:"omitempty,url"`
	ParentName           *string                       `json:"parentName"`
	ParentPhoneNumber    *string                       `json:"parentPhoneNumber"`
	PaymentStyle         PaymentStyle                  `json:"paymentStyle" validate:"payment_style"`
	PaymentType          PaymentType                   `json:"paymentType" validate:"payment_type"`
	PaymentTypeValue     *float64                      `json:"paymentTypeValue"`
	SettlementStyle      SettlementStyle               `json:"settlementStyle" validate:"settlement_style"`
	SettlementStyleValue *int                          `json:"settlementStyleValue"`
	SettlementStyleDay   *int                          `json:"settlementStyleDay"`
	HouseAddress         *string                       `json:"houseAddress"`
	HouseIdentifier      *string                       `json:"hoseInfo"`
	HouseCoordinate      *models.Coordinate            `json:"houseCoordinate"`
	Routine              []CreateStudentRoutinePlanDTO `json:"routine" validate:"required"`
}

type CreateStudentRoutinePlanDTO struct {
	WeekDay   models.WeekDay `json:"weekDay" validate:"required,weekday"`
	StartHour int            `json:"startHour" validate:"required"`
	Duration  int            `json:"duration" validate:"required"`
	Price     float64        `json:"price" validate:"required"`
}

type UpdateStudentDTO struct {
	IDUser               int                           `json:"idUser"`
	ID                   int                           `json:"id"`
	Name                 string                        `json:"name" validate:"required"`
	BirthDay             *string                       `json:"birthDay" validate:"omitempty,datetime"`
	DisplayColor         string                        `json:"displayColor" validate:"required,hexcolor"`
	Picture              *string                       `json:"picture" validate:"omitempty,url"`
	Gender               *models.Gender                `json:"gender" validate:"omitempty,gender"`
	ParentName           *string                       `json:"parentName"`
	ParentPhoneNumber    *string                       `json:"parentPhoneNumber"`
	PaymentStyle         PaymentStyle                  `json:"paymentStyle" validate:"payment_style"`
	PaymentType          PaymentType                   `json:"paymentType" validate:"payment_type"`
	PaymentTypeValue     *float64                      `json:"paymentTypeValue"`
	SettlementStyle      SettlementStyle               `json:"settlementStyle" validate:"settlement_style"`
	SettlementStyleValue *int                          `json:"settlementStyleValue"`
	SettlementStyleDay   *int                          `json:"settlementStyleDay"`
	HouseAddress         *string                       `json:"houseAddress"`
	HouseIdentifier      *string                       `json:"hoseInfo"`
	HouseCoordinate      *models.Coordinate            `json:"houseCoordinate"`
	Routine              []UpdateStudentRoutinePlanDTO `json:"routine" validate:"required"`
}

type UpdateStudentRoutinePlanDTO struct {
	ID        *int            `json:"id"`
	WeekDay   *models.WeekDay `json:"weekDay" validate:"required_ifid,weekday_ifid"`
	StartHour *int            `json:"startHour" validate:"required_ifid"`
	Duration  *int            `json:"duration" validate:"required_ifid"`
	Price     *float64        `json:"price" validate:"required_ifid"`
}

type DeleteStudentDTO struct {
	IDUser int `json:"idUser"`
	ID     int `json:"id"`
}
