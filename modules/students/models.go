package students

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/models"
)

type Student struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	BirthDay          *string            `json:"birthDay"`
	DisplayColor      string             `json:"displayColor"`
	Picture           *string            `json:"picture"`
	ParentName        *string            `json:"parentName"`
	ParentPhoneNumber *string            `json:"parentPhoneNumber"`
	HouseAddress      *string            `json:"houseAddress"`
	HouseIdentifier   *string            `json:"hoseIdentifier"`
	HouseCoordinate   *models.Coordinate `json:"houseCoordinate"`
	BasePrice         float64            `json:"basePrice"`
	Routine           []RoutinePlan      `json:"routine"`
	IsDeleted         bool               `json:"isDeleted"`
	CreatedAt         time.Time          `json:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt"`
}

func (s *Student) ToStudentEntity() studentEntity {
	var latitude *float64
	var longitude *float64
	if s.HouseCoordinate != nil {
		latitude = &s.HouseCoordinate.Latitude
		longitude = &s.HouseCoordinate.Longitude
	}
	return studentEntity{
		ID:                       s.ID,
		Name:                     s.Name,
		BirthDay:                 s.BirthDay,
		DisplayColor:             s.DisplayColor,
		Picture:                  s.Picture,
		ParentName:               s.ParentName,
		ParentPhoneNumber:        s.ParentPhoneNumber,
		HouseAddress:             s.HouseAddress,
		HouseIdentifier:          s.HouseIdentifier,
		HouseCoordinateLatitude:  latitude,
		HouseCoordinateLongitude: longitude,
		BasePrice:                s.BasePrice,
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
		IdStudent: idStudent,
		WeekDay:   r.WeekDay,
		StartHour: r.StartHour,
		Duration:  r.Duration,
		Price:     r.Price,
		IsDeleted: r.IsDeleted,
		CreatedAt: r.CreatedAt,
	}
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

type ScheduleDay struct {
	Day          Day           `json:"day"`
	Appointments []Appointment `json:"appointments"`
	Routines     []RoutinePlan `json:"routines"`
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
