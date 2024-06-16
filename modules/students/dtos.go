package students

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type CreateStudentRoutineDTO struct {
	WeekDay   models.WeekDay `json:"weekDay" validate:"required,weekday"`
	StartHour int            `json:"startHour" validate:"required"`
	Duration  int            `json:"duration" validate:"required"`
	Price     *float64       `json:"price" validate:"required"`
}

type CreateStudentDTO struct {
	Name              string                    `json:"name" validate:"required"`
	BirthDay          *string                   `json:"birthDay" validate:"omitempty,datetime"`
	DisplayColor      string                    `json:"displayColor" validate:"omitempty,hexcolor"`
	Picture           *string                   `json:"picture" validate:"omitempty,url"`
	ParentName        *string                   `json:"parentName"`
	ParentPhoneNumber *string                   `json:"parentPhoneNumber"`
	HouseAddress      *string                   `json:"houseAddress"`
	HouseIdentifier   *string                   `json:"hoseInfo"`
	HouseCoordinate   *models.Coordinate        `json:"houseCoordinate"`
	BasePrice         float64                   `json:"basePrice" validate:"required"`
	Routine           []CreateStudentRoutineDTO `json:"routine" validate:"required"`
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
		Routine: utils.Map(c.Routine, func(r CreateStudentRoutineDTO) RoutinePlan {
			var price = r.Price
			if price == nil {
				price = &c.BasePrice
			}

			return RoutinePlan{
				WeekDay:   r.WeekDay,
				StartHour: r.StartHour,
				Duration:  r.Duration,
				Price:     *price,
			}
		}),
	}
}

type UpdateStudentRoutineDTO struct {
	ID        *int            `json:"id"`
	WeekDay   *models.WeekDay `json:"weekDay" validate:"required_ifid,weekday_ifid"`
	StartHour *int            `json:"startHour" validate:"required_ifid"`
	Duration  *int            `json:"duration" validate:"required_ifid"`
	Price     *float64        `json:"price" validate:"required_ifid"`
}

func (c *UpdateStudentRoutineDTO) ToRoutinePlanEntity(idStudent int) routinePlanEntity {
	if c.ID == nil {
		return routinePlanEntity{
			ID:        nil,
			IdStudent: idStudent,
			WeekDay:   *c.WeekDay,
			StartHour: *c.StartHour,
			Duration:  *c.Duration,
			Price:     *c.Price,
		}
	}
	return routinePlanEntity{
		ID: c.ID,
	}
}

type UpdateStudentDTO struct {
	CreateStudentDTO
	ID      int                       `json:"id" validate:"required"`
	Routine []UpdateStudentRoutineDTO `json:"routine" validate:"required"`
}

func (c *UpdateStudentDTO) ToStudentEntity() studentEntity {
	var latitude *float64
	var longitude *float64
	if c.HouseCoordinate != nil {
		latitude = &c.HouseCoordinate.Latitude
		longitude = &c.HouseCoordinate.Longitude

	}
	return studentEntity{
		ID:                       c.ID,
		Name:                     c.Name,
		BirthDay:                 c.BirthDay,
		DisplayColor:             c.DisplayColor,
		Picture:                  c.Picture,
		ParentName:               c.ParentName,
		ParentPhoneNumber:        c.ParentPhoneNumber,
		HouseAddress:             c.HouseAddress,
		HouseIdentifier:          c.HouseIdentifier,
		HouseCoordinateLatitude:  latitude,
		HouseCoordinateLongitude: longitude,
		BasePrice:                c.BasePrice,
		UpdatedAt:                time.Now(),
	}
}
