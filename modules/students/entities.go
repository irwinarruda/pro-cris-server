package students

import "time"

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Student struct {
	Id                int         `json:"id"`
	Name              string      `json:"name"`
	BirthDay          *string     `json:"birthDay"`
	DisplayColor      string      `json:"displayColor"`
	Picture           *string     `json:"picture"`
	ParentName        *string     `json:"parentName"`
	ParentPhoneNumber *string     `json:"parentPhoneNumber"`
	HouseAddress      *string     `json:"houseAddress"`
	HouseIdentifier   *string     `json:"hoseInfo"`
	HouseCoordinate   *Coordinate `json:"houseCoordinate"`
	BasePrice         float64     `json:"basePrice"`
	IsDeleted         bool        `json:"isDeleted"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}

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

type Day struct {
	Id                int    `json:"id"`
	Day               string `json:"day"`
	Month             string `json:"month"`
	Year              string `json:"year"`
	HasRoutineStarted bool   `json:"hasRoutineStarted"`
}

type Appointment struct {
	Id         int       `json:"id"`
	IdStudent  int       `json:"idStudent"`
	Day        Day       `json:"day"`
	StartHour  string    `json:"startHour"`
	Duration   int       `json:"duration"`
	Price      float64   `json:"price"`
	IsSettled  bool      `json:"isSettled"`
	IsPrePaid  bool      `json:"isPrePaid"`
	IsCanceled bool      `json:"isCanceled"`
	IsDeleted  bool      `json:"isDeleted"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Holiday struct {
	Day  Day    `json:"day"`
	Name string `json:"name"`
}

type Routine struct {
	IdStudent string  `json:"idStudent"`
	WeekDay   WeekDay `json:"weekDay"`
	StartHour string  `json:"startHour"`
	Duration  int     `json:"duration"` // milisseconds
	Price     float64 `json:"price"`
}

type ScheduleDay struct {
	Day          Day           `json:"day"`
	Appointments []Appointment `json:"appointments"`
	Routines     []Routine     `json:"routines"`
	Holidays     []Holiday     `json:"holidays"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Picture   string    `json:"picture"`
	Provider  string    `json:"provider"` // Google
	IsDeleted bool      `json:"isDeleted"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}