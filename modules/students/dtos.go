package students

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/entities"
)

type CreateStudentRoutineDTO struct {
	WeekDay   entities.WeekDay `json:"weekDay" validate:"required,weekday"`
	StartHour string           `json:"startHour" validate:"required"`
	Duration  int              `json:"duration" validate:"required"`
	Price     float64          `json:"price" validate:"required"`
}

type CreateStudentDTO struct {
	Name              string               `json:"name" validate:"required"`
	BirthDay          *string              `json:"birthDay" validate:"omitempty,datetime"`
	DisplayColor      string               `json:"displayColor" validate:"omitempty,hexcolor"`
	Picture           *string              `json:"picture" validate:"omitempty,url"`
	ParentName        *string              `json:"parentName"`
	ParentPhoneNumber *string              `json:"parentPhoneNumber"`
	HouseAddress      *string              `json:"houseAddress"`
	HouseIdentifier   *string              `json:"hoseInfo"`
	HouseCoordinate   *entities.Coordinate `json:"houseCoordinate"`
	BasePrice         float64              `json:"basePrice" validate:"required"`
	Routine           []Routine            `json:"routine" validate:"required"`
}

func (c *CreateStudentDTO) ToStudent() Student {
	return Student{
		Name:              c.Name,
		BirthDay:          c.BirthDay,
		DisplayColor:      c.DisplayColor,
		Picture:           c.Picture,
		ParentName:        c.ParentName,
		ParentPhoneNumber: c.ParentPhoneNumber,
		HouseAddress:      c.HouseAddress,
		HouseIdentifier:   c.HouseIdentifier,
		HouseCoordinate:   c.HouseCoordinate,
		BasePrice:         c.BasePrice,
		IsDeleted:         false,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}
