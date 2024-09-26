package appointments

import (
	"time"
)

type Appointment struct {
	ID          int                `json:"id"`
	CalendarDay time.Time          `json:"calendarDay"`
	StartHour   int                `json:"startHour"` // milisseconds
	Duration    int                `json:"duration"`  // milisseconds
	Price       float64            `json:"price"`
	IsExtra     bool               `json:"isExtra"`
	IsPaid      bool               `json:"isPaid"`
	Student     AppointmentStudent `json:"student"`
	IsDeleted   bool               `json:"isDeleted"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	// IsSettled   bool               `json:"isSettled"`
}

type AppointmentStudent struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	DisplayColor string  `json:"displayColor"`
	Picture      *string `json:"picture"`
}

type CalendarDay struct {
	Id    int `json:"id"`
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Holiday struct {
	Day  CalendarDay `json:"day"`
	Name string      `json:"name"`
}
