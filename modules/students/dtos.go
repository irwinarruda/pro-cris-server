package students

import "github.com/irwinarruda/pro-cris-server/shared/models"

type CreateStudentDTO struct {
	Name              string                        `json:"name" validate:"required"`
	BirthDay          *string                       `json:"birthDay" validate:"omitempty,datetime"`
	DisplayColor      string                        `json:"displayColor" validate:"omitempty,hexcolor"`
	Picture           *string                       `json:"picture" validate:"omitempty,url"`
	ParentName        *string                       `json:"parentName"`
	ParentPhoneNumber *string                       `json:"parentPhoneNumber"`
	HouseAddress      *string                       `json:"houseAddress"`
	HouseIdentifier   *string                       `json:"hoseInfo"`
	HouseCoordinate   *models.Coordinate            `json:"houseCoordinate"`
	BasePrice         float64                       `json:"basePrice" validate:"required"`
	Routine           []CreateStudentRoutinePlanDTO `json:"routine" validate:"required"`
}

type CreateStudentRoutinePlanDTO struct {
	WeekDay   models.WeekDay `json:"weekDay" validate:"required,weekday"`
	StartHour int            `json:"startHour" validate:"required"`
	Duration  int            `json:"duration" validate:"required"`
	Price     *float64       `json:"price"`
}

type UpdateStudentDTO struct {
	ID                int                           `json:"id"`
	Name              string                        `json:"name" validate:"required"`
	BirthDay          *string                       `json:"birthDay" validate:"omitempty,datetime"`
	DisplayColor      string                        `json:"displayColor" validate:"omitempty,hexcolor"`
	Picture           *string                       `json:"picture" validate:"omitempty,url"`
	ParentName        *string                       `json:"parentName"`
	ParentPhoneNumber *string                       `json:"parentPhoneNumber"`
	HouseAddress      *string                       `json:"houseAddress"`
	HouseIdentifier   *string                       `json:"hoseInfo"`
	HouseCoordinate   *models.Coordinate            `json:"houseCoordinate"`
	BasePrice         float64                       `json:"basePrice" validate:"required"`
	Routine           []UpdateStudentRoutinePlanDTO `json:"routine" validate:"required"`
}

type UpdateStudentRoutinePlanDTO struct {
	ID        *int            `json:"id"`
	WeekDay   *models.WeekDay `json:"weekDay" validate:"required_ifid,weekday_ifid"`
	StartHour *int            `json:"startHour" validate:"required_ifid"`
	Duration  *int            `json:"duration" validate:"required_ifid"`
	Price     *float64        `json:"price" validate:"required_ifid"`
}
