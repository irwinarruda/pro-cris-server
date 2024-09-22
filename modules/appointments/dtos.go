package appointments

import "time"

type GetAppointmentDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
	ID        int `json:"id" validate:"required"`
}

type CreateAppointmentDTO struct {
	IDAccount   int       `json:"idAccount" validate:"required"`
	IDStudent   int       `json:"idStudent" validate:"required"`
	CalendarDay time.Time `json:"calendarDay"`
	StartHour   string    `json:"startHour"`
	Duration    int       `json:"duration"`
	Price       float64   `json:"price"`
	IsExtra     bool      `json:"isExtra"`
	IsPaid      bool      `json:"isPaid"`
}

type CreateDailyAppointmentsByStudentsRoutineDTO struct {
	IDAccount   int       `json:"idAccount" validate:"required"`
	CalendarDay time.Time `json:"calendarDay"`
}

type UpdateAppointmentDTO struct {
	IDAccount int     `json:"idAccount" validate:"required"`
	ID        int     `json:"id" validate:"required"`
	Price     float64 `json:"price"`
	IsExtra   bool    `json:"isExtra"`
	IsPaid    bool    `json:"isPaid"`
}

type DeleteAppointmentDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
	ID        int `json:"id" validate:"required"`
}
