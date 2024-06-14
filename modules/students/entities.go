package students

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/entities"
)

type Student struct {
	ID                       int                  `json:"id"`
	Name                     string               `json:"name"`
	BirthDay                 *string              `json:"birthDay"`
	DisplayColor             string               `json:"displayColor"`
	Picture                  *string              `json:"picture"`
	ParentName               *string              `json:"parentName"`
	ParentPhoneNumber        *string              `json:"parentPhoneNumber"`
	HouseAddress             *string              `json:"houseAddress"`
	HouseIdentifier          *string              `json:"hoseIdentifier"`
	HouseCoordinate          *entities.Coordinate `json:"houseCoordinate"`
	BasePrice                float64              `json:"basePrice"`
	Routine                  []Routine            `json:"routine"`
	IsDeleted                bool                 `json:"isDeleted"`
	CreatedAt                time.Time            `json:"createdAt"`
	UpdatedAt                time.Time            `json:"updatedAt"`
	HouseCoordinateLatitude  *float64             `json:"houseCoordinateLatitude,omitempty"`
	HouseCoordinateLongitude *float64             `json:"houseCoordinateLongitude,omitempty"`
}

type Day struct {
	Id                int    `json:"id"`
	Day               string `json:"day"`
	Month             string `json:"month"`
	Year              string `json:"year"`
	HasRoutineStarted bool   `json:"hasRoutineStarted"`
}

type Appointment struct {
	Id         int       `json:"id"`
	IdStudent  int       `json:"idStudent"`
	Day        Day       `json:"day"`
	StartHour  string    `json:"startHour"`
	Duration   int       `json:"duration"`
	Price      float64   `json:"price"`
	IsSettled  bool      `json:"isSettled"`
	IsPrePaid  bool      `json:"isPrePaid"`
	IsCanceled bool      `json:"isCanceled"`
	IsDeleted  bool      `json:"isDeleted"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Holiday struct {
	Day  Day    `json:"day"`
	Name string `json:"name"`
}

type Routine struct {
	IdStudent string           `json:"idStudent"`
	WeekDay   entities.WeekDay `json:"weekDay"`
	StartHour string           `json:"startHour"`
	Duration  int              `json:"duration"` // milisseconds
	Price     float64          `json:"price"`
}

type ScheduleDay struct {
	Day          Day           `json:"day"`
	Appointments []Appointment `json:"appointments"`
	Routines     []Routine     `json:"routines"`
	Holidays     []Holiday     `json:"holidays"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Picture   string    `json:"picture"`
	Provider  string    `json:"provider"` // Google
	IsDeleted bool      `json:"isDeleted"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
