package date

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/models"
)

type DateService struct{}

type IDateService = *DateService

func NewDateService() *DateService {
	return proinject.Resolve(&DateService{})
}

func (d *DateService) GetWeekDayFromDate(date time.Time) models.WeekDay {
	return models.WeekDay(date.Weekday().String())
}
