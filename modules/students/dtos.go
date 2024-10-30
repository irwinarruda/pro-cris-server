package students

import "github.com/irwinarruda/pro-cris-server/shared/models"

type GetAllStudentsDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
}

type GetStudentDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
	ID        int `json:"id" validate:"required"`
}

type DoesStudentExistsDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
	ID        int `json:"id" validate:"required"`
}

type CreateStudentDTO struct {
	CreateStudentSettlementOptionsDTO
	IDAccount         int                           `json:"idAccount" validate:"required"`
	Name              string                        `json:"name" validate:"required"`
	BirthDay          *string                       `json:"birthDay" validate:"omitempty,datetime=2006-01-02"`
	DisplayColor      string                        `json:"displayColor" validate:"omitempty,hexcolor"`
	Gender            *models.Gender                `json:"gender" validate:"gender"`
	Picture           *string                       `json:"picture" validate:"omitempty,url"`
	ParentName        *string                       `json:"parentName"`
	ParentPhoneNumber *string                       `json:"parentPhoneNumber"`
	HouseAddress      *string                       `json:"houseAddress"`
	HouseIdentifier   *string                       `json:"hoseInfo"`
	HouseCoordinate   *models.Coordinate            `json:"houseCoordinate"`
	Routine           []CreateStudentRoutinePlanDTO `json:"routine" validate:"required,dive"`
}

type CreateStudentSettlementOptionsDTO struct {
	PaymentStyle         models.PaymentStyle    `json:"paymentStyle" validate:"payment_style"`
	PaymentType          models.PaymentType     `json:"paymentType" validate:"payment_type"`
	PaymentTypeValue     *float64               `json:"paymentTypeValue"`
	SettlementStyle      models.SettlementStyle `json:"settlementStyle" validate:"settlement_style"`
	SettlementStyleValue *int                   `json:"settlementStyleValue"`
	SettlementStyleDay   *int                   `json:"settlementStyleDay"`
}

type CreateStudentRoutinePlanDTO struct {
	WeekDay   models.WeekDay `json:"weekDay" validate:"required,weekday"`
	StartHour int            `json:"startHour" validate:"required,min=0,max=86400000"`
	Duration  int            `json:"duration" validate:"required,min=0,max=86400000"`
	Price     float64        `json:"price" validate:"required,min=0"`
}

type UpdateStudentDTO struct {
	UpdateStudentSettlementOptionsDTO
	IDAccount         int                           `json:"idAccount" validate:"required"`
	ID                int                           `json:"id" validate:"required"`
	Name              string                        `json:"name" validate:"required"`
	BirthDay          *string                       `json:"birthDay" validate:"omitempty,datetime=2006-01-02"`
	DisplayColor      string                        `json:"displayColor" validate:"required,hexcolor"`
	Picture           *string                       `json:"picture" validate:"omitempty,url"`
	Gender            *models.Gender                `json:"gender" validate:"gender"`
	ParentName        *string                       `json:"parentName"`
	ParentPhoneNumber *string                       `json:"parentPhoneNumber"`
	HouseAddress      *string                       `json:"houseAddress"`
	HouseIdentifier   *string                       `json:"hoseInfo"`
	HouseCoordinate   *models.Coordinate            `json:"houseCoordinate"`
	Routine           []UpdateStudentRoutinePlanDTO `json:"routine" validate:"required,dive"`
}

type UpdateStudentSettlementOptionsDTO struct {
	PaymentStyle         models.PaymentStyle    `json:"paymentStyle" validate:"payment_style"`
	PaymentType          models.PaymentType     `json:"paymentType" validate:"payment_type"`
	PaymentTypeValue     *float64               `json:"paymentTypeValue"`
	SettlementStyle      models.SettlementStyle `json:"settlementStyle" validate:"settlement_style"`
	SettlementStyleValue *int                   `json:"settlementStyleValue"`
	SettlementStyleDay   *int                   `json:"settlementStyleDay"`
}

type UpdateStudentRoutinePlanDTO struct {
	ID        *int            `json:"id"`
	WeekDay   *models.WeekDay `json:"weekDay" validate:"omitnil,required"`
	StartHour *int            `json:"startHour" validate:"omitnil,min=0,max=86400000"`
	Duration  *int            `json:"duration" validate:"omitnil,min=0,max=86400000"`
	Price     *float64        `json:"price" validate:"omitnil,min=0"`
}

type DeleteStudentDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
	ID        int `json:"id" validate:"required"`
}
