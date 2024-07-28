package appointments

type GetAppointmentDTO struct {
	IDAccount int `json:"idAccount"`
	ID        int `json:"id"`
}

type CreateAppointmentDTO struct {
	IDAccount   int                             `json:"idAccount"`
	IDStudent   int                             `json:"idStudent"`
	CalendarDay CreateAppointmentCalendarDayDTO `json:"calendarDay"`
	StartHour   string                          `json:"startHour"`
	Duration    int                             `json:"duration"`
	Price       float64                         `json:"price"`
	IsExtra     bool                            `json:"isExtra"`
	IsPaid      bool                            `json:"isPaid"`
}

type CreateAppointmentCalendarDayDTO struct {
	ID    int `json:"id"`
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
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
