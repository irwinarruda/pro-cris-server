package appointments

type CreateAppointmentDTO struct {
	IDStudent   int                             `json:"idStudent"`
	CalendarDay CreateAppointmentCalendarDayDTO `json:"calendarDay"`
	StartHour   string                          `json:"startHour"`
	Duration    int                             `json:"duration"`
	Price       float64                         `json:"price"`
	IsExtra     bool                            `json:"isExtra"`
}

type CreateAppointmentCalendarDayDTO struct {
	ID    int `json:"id"`
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}
