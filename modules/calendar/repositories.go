package calendar

type ICalendarRepository interface {
	CreateCalendarDay(day, month, year int) (int, error)
	GetCalendarDayByDate(day, month, year int) (CalendarDay, error)
	GetCalendarDayByID(id int) (CalendarDay, error)
	ResetCalendarDays()
}
