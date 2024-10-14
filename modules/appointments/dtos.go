package appointments

import "time"

type GetAppointmentDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
	ID        int `json:"id" validate:"required"`
}

type GetAppointmentsByDateRangeDTO struct {
	IDAccount   int       `json:"idAccount" validate:"required"`
	IDStudent   int       `json:"idStudent" validate:"required"`
	InitialDate time.Time `json:"initialDate" validate:"required"`
	FinalDate   time.Time `json:"finalDate" validate:"required"`
}

type CreateAppointmentDTO struct {
	IDAccount   int       `json:"idAccount" validate:"required"`
	IDStudent   int       `json:"idStudent" validate:"required"`
	CalendarDay time.Time `json:"calendarDay" validate:"required"`
	StartHour   int       `json:"startHour" validate:"required"` // milisseconds
	Duration    int       `json:"duration" validate:"required"`  // milisseconds
	Price       float64   `json:"price" validate:"required"`
	IsExtra     bool      `json:"isExtra"`
	IsPaid      bool      `json:"isPaid"`
}

type CreateDailyAppointmentsByStudentsRoutineDTO struct {
	IDAccount   int       `json:"idAccount" validate:"required"`
	CalendarDay time.Time `json:"calendarDay" validate:"required"`
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
