package models

import (
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

func (w WeekDay) String() string {
	return string(w)
}

func (w *WeekDay) UnmarshalJSON(b []byte) (err error) {
	return utils.UnmarshalEnum(w, GetWeekDaysString(), b)
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
