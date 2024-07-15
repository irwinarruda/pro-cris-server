package calendar

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type CalendarService struct {
	CalendarRepository ICalendarRepository `inject:"calendar_repository"`
}

func NewCalendarService() *CalendarService {
	return proinject.Resolve(&CalendarService{})
}

func (a *CalendarService) CreateCalendarDayIfNotExists(day, month, year int) (int, error) {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	if date.Day() != day || date.Month() != time.Month(month) || date.Year() != year {
		return 0, utils.NewAppError("Invalid date.", true, nil)
	}
	calendarDay, err := a.CalendarRepository.GetCalendarDayByDate(day, month, year)
	if err == nil {
		return calendarDay.ID, nil
	}
	id, err := a.CalendarRepository.CreateCalendarDay(day, month, year)
	if err != nil {
		return 0, utils.NewAppError("Error creating a date.", false, nil)
	}
	return id, nil
}

func (a *CalendarService) GetCalendarDayByID(id int) (CalendarDay, error) {
	return a.CalendarRepository.GetCalendarDayByID(id)
}
