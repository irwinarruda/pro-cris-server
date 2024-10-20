package models

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type WeekDay string

const (
	Monday    WeekDay = "Monday"
	Tuesday   WeekDay = "Tuesday"
	Wednesday WeekDay = "Wednesday"
	Thursday  WeekDay = "Thursday"
	Friday    WeekDay = "Friday"
	Saturday  WeekDay = "Saturday"
	Sunday    WeekDay = "Sunday"
)

func (w WeekDay) ToInt() int {
	switch w {
	case Sunday:
		return 0
	case Monday:
		return 1
	case Tuesday:
		return 2
	case Wednesday:
		return 3
	case Thursday:
		return 4
	case Friday:
		return 5
	case Saturday:
		return 6
	}
	return 0
}

func (w WeekDay) FromInt(i int) WeekDay {
	return WeekDay(time.Weekday(i).String())
}

func (w WeekDay) Before() WeekDay {
	intWeekDay := w.ToInt()
	if intWeekDay == 0 {
		return w.FromInt(6)
	}
	return w.FromInt(intWeekDay - 1)
}

func (w WeekDay) After() WeekDay {
	intWeekDay := w.ToInt()
	if intWeekDay == 6 {
		return w.FromInt(0)
	}
	return w.FromInt(intWeekDay + 1)
}

func (w WeekDay) String() string {
	return string(w)
}

func (w *WeekDay) UnmarshalJSON(b []byte) (err error) {
	return utils.UnmarshalEnum(w, GetWeekDaysString(), b)
}

func FromTime(i time.Time) WeekDay {
	return WeekDay(i.Weekday().String())
}

func GetWeekDaysString() []string {
	return []string{
		Monday.String(),
		Tuesday.String(),
		Wednesday.String(),
		Thursday.String(),
		Friday.String(),
		Saturday.String(),
		Sunday.String(),
	}
}
