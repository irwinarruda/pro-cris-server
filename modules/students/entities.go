package students

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/models"
)

type Student struct {
	models.SettlementOptions
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	BirthDay          *string            `json:"birthDay"`
	DisplayColor      string             `json:"displayColor"`
	Picture           *string            `json:"picture"`
	Gender            *models.Gender     `json:"gender"`
	ParentName        *string            `json:"parentName"`
	ParentPhoneNumber *string            `json:"parentPhoneNumber"`
	HouseAddress      *string            `json:"houseAddress"`
	HouseIdentifier   *string            `json:"hoseIdentifier"`
	HouseCoordinate   *models.Coordinate `json:"houseCoordinate"`
	Routine           []RoutinePlan      `json:"routine"`
	IsDeleted         bool               `json:"isDeleted"`
	CreatedAt         time.Time          `json:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt"`
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
