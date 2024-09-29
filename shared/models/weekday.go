package models

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

func ToWeekDay(day any) WeekDay {
	return WeekDay(day.(string))
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
