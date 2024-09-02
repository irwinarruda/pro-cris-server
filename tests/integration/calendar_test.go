package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/calendar"
	calendarresources "github.com/irwinarruda/pro-cris-server/modules/calendar/resources"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/stretchr/testify/assert"
)

func TestCalendarServiceHappyPath(t *testing.T) {
	beforeEachCalendar()
	var assert = assert.New(t)
	var calendarService = calendar.NewCalendarService()

	id1, _ := calendarService.CreateCalendarDayIfNotExists(1, 1, 2024)
	calendarDay, err := calendarService.GetCalendarDayByID(id1)
	assert.NoError(err, "Should return get the created calendarDay.")
	assert.Equal(1, calendarDay.Day, "Should return Price.")
	assert.Equal(1, calendarDay.Month, "Should return Duration.")
	assert.Equal(2024, calendarDay.Year, "Should return IDStudent.")
	id2, _ := calendarService.CreateCalendarDayIfNotExists(1, 1, 2024)
	assert.Equal(id1, id2, "Should return the same ID for the same date.")
}

func TestCalendarServiceErrorPath(t *testing.T) {
	beforeEachCalendar()

	var assert = assert.New(t)
	var calendarService = calendar.NewCalendarService()

	_, err := calendarService.GetCalendarDayByID(10)
	assert.Error(err, "Should not return an unexistent calendar.")

	_, err = calendarService.CreateCalendarDayIfNotExists(30, 2, 2024)
	assert.Error(err, "Should return an error with invalid date.")
}

func beforeEachCalendar() {
	proinject.Register("env", configs.GetEnv("../../.env"))
	proinject.Register("db", configs.GetDb())

	calendarRepository := calendarresources.NewDbCalendarRepository()
	calendarRepository.ResetCalendarDays()
	proinject.Register("calendar_repository", calendarRepository)
}
