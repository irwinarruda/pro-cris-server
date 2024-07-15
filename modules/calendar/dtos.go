package calendar

type CreateCalendarDayDTO struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}
