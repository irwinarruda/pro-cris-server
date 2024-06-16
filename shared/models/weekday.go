package models

type WeekDay = string

const (
	Monday    WeekDay = "Monday"
	Tuesday   WeekDay = "Tuesday"
	Wednesday WeekDay = "Wednesday"
	Thursday  WeekDay = "Thursday"
	Friday    WeekDay = "Friday"
	Saturday  WeekDay = "Saturday"
	Sunday    WeekDay = "Sunday"
)

func GetWeekDays() []WeekDay {
	return []WeekDay{
		Monday,
		Tuesday,
		Wednesday,
		Thursday,
		Friday,
		Saturday,
		Sunday,
	}
}
