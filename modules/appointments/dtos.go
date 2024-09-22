package appointments

import "time"

type GetAppointmentDTO struct {
	IDAccount int `json:"idAccount"`
	ID        int `json:"id"`
}

type CreateAppointmentDTO struct {
	IDAccount   int       `json:"idAccount"`
	IDStudent   int       `json:"idStudent"`
	CalendarDay time.Time `json:"calendarDay"`
	StartHour   string    `json:"startHour"`
	Duration    int       `json:"duration"`
	Price       float64   `json:"price"`
	IsExtra     bool      `json:"isExtra"`
	IsPaid      bool      `json:"isPaid"`
}

type CreateDailyAppointmentsByStudentsRoutineDTO struct {
	IDAccount   int       `json:"idAccount"`
	CalendarDay time.Time `json:"calendarDay"`
}

type UpdateAppointmentDTO struct {
	IDAccount int     `json:"idAccount"`
	ID        int     `json:"id"`
	Price     float64 `json:"price"`
	IsExtra   bool    `json:"isExtra"`
	IsPaid    bool    `json:"isPaid"`
}

type DeleteAppointmentDTO struct {
	IDAccount int `json:"idAccount"`
	ID        int `json:"id"`
}
