package auth

type User struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Email           string           `json:"email"`
	Picutre         *string          `json:"picture"`
	EmailVerified   bool             `json:"emailVerified"`
	Permissions     []Permission     `json:"permissions"`
	TeacherFeatures *TeacherFeatures `json:"teacherFeatures"`
}

type TeacherFeatures struct {
	ID                            int  `json:"id"`
	MaxStudents                   int  `json:"maxStudents"`
	CanGenerateStudentReport      bool `json:"canGenerateStudentReport"`
	CanManuallyCreateAppointments bool `json:"canManuallyCreateAppointments"`
}

type Permission = string

const (
	Admin   Permission = "Admin"
	Teacher Permission = "Teacher"
)
