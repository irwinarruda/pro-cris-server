package appointments

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/modules/calendar"
)

type Appointment struct {
	ID          int                  `json:"id"`
	CalendarDay calendar.CalendarDay `json:"calendarDay"`
	StartHour   string               `json:"startHour"`
	Duration    int                  `json:"duration"`
	Price       float64              `json:"price"`
	IsExtra     bool                 `json:"isExtra"`
	Student     AppointmentStudent   `json:"student"`
	IsDeleted   bool                 `json:"isDeleted"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
	// IsSettled   bool               `json:"isSettled"`
	// IsPaid      bool               `json:"isPaid"`
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
