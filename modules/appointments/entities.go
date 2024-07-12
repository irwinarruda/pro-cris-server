package appointments

import (
	"time"
)

type Appointment struct {
	ID         int                `json:"id"`
	Day        Day                `json:"day"`
	StartHour  string             `json:"startHour"`
	Duration   int                `json:"duration"`
	Price      float64            `json:"price"`
	IsSettled  bool               `json:"isSettled"`
	IsPaid     bool               `json:"isPrePaid"`
	IsCanceled bool               `json:"isCanceled"`
	Student    AppointmentStudent `json:"student"`
	IsDeleted  bool               `json:"isDeleted"`
	CreatedAt  time.Time          `json:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt"`
}

type AppointmentStudent struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	DisplayColor string  `json:"displayColor"`
	Picture      *string `json:"picture"`
}

type Day struct {
	Id                int    `json:"id"`
	Day               string `json:"day"`
	Month             string `json:"month"`
	Year              string `json:"year"`
	HasRoutineStarted bool   `json:"hasRoutineStarted"`
}

type Holiday struct {
	Day  Day    `json:"day"`
	Name string `json:"name"`
}

// type ScheduleDay struct {
// 	Day          Day           `json:"day"`
// 	Appointments []Appointment `json:"appointments"`
// 	Routines     []RoutinePlan `json:"routines"`
// 	Holidays     []Holiday     `json:"holidays"`
// }
