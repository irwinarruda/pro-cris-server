package calendarresources

import (
	"fmt"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/calendar"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type DbCalendarRepository struct {
	Db configs.Db `inject:"db"`
}

func NewDbCalendarRepository() *DbCalendarRepository {
	return proinject.Resolve(&DbCalendarRepository{})
}

func (a *DbCalendarRepository) CreateCalendarDay(day, month, year int) (int, error) {
	id := 0
	sql := fmt.Sprintf(`
    INSERT INTO "calendar_day"(
      day,
      month,
      year
    ) %s
    RETURNING id;
  `, utils.SqlValues(1, 3))
	err := a.Db.Raw(sql, day, month, year).Scan(&id).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *DbCalendarRepository) GetCalendarDayByID(id int) (calendar.CalendarDay, error) {
	calendarDayE := []calendar.CalendarDay{}
	sql := `SELECT * FROM "calendar_day" WHERE id = ?;`
	err := a.Db.Raw(sql, id).Scan(&calendarDayE).Error
	if err != nil {
		return calendar.CalendarDay{}, err
	}
	if len(calendarDayE) == 0 {
		return calendar.CalendarDay{}, utils.NewAppError("Calendar day not found.", true, nil)
	}
	return calendarDayE[0], nil
}

func (a *DbCalendarRepository) GetCalendarDayByDate(day int, month int, year int) (calendar.CalendarDay, error) {
	calendarDayE := []calendar.CalendarDay{}
	sql := `SELECT * FROM "calendar_day" WHERE day = ? AND month = ? AND year = ?;`
	err := a.Db.Raw(sql, day, month, year).Scan(&calendarDayE).Error
	if err != nil {
		return calendar.CalendarDay{}, err
	}
	if len(calendarDayE) == 0 {
		return calendar.CalendarDay{}, utils.NewAppError("Calendar day not found.", true, nil)
	}
	return calendarDayE[0], nil
}

func (a *DbCalendarRepository) ResetCalendarDays() {
	a.Db.Exec(`DELETE FROM "calendar_day";`)
	a.Db.Exec(`ALTER SEQUENCE calendar_day_id_seq RESTART WITH 1;`)
}
