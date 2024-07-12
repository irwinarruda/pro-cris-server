package auth

import "time"

type Account struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Email           string           `json:"email"`
	Picture         *string          `json:"picture"`
	EmailVerified   bool             `json:"emailVerified"`
	Provider        LoginProvider    `json:"provider"`
	Permissions     []Permission     `json:"permissions"`
	TeacherFeatures *TeacherFeatures `json:"teacherFeatures"`
	IsDeleted       bool             `json:"isDeleted"`
	CreatedAt       time.Time        `json:"createdAt"`
	UpdatedAt       time.Time        `json:"updatedAt"`
}

type TeacherFeatures struct {
	ID                            int  `json:"id"`
	MaxStudents                   int  `json:"maxStudents"`
	CanGenerateStudentReport      bool `json:"canGenerateStudentReport"`
	CanManuallyCreateAppointments bool `json:"canManuallyCreateAppointments"`
}

type Permission = string

const (
	PermissionAdmin   Permission = "Admin"
	PermissionTeacher Permission = "Teacher"
)

type LoginProvider = string

const (
	LoginProviderGoogle LoginProvider = "Google"
)

func GetLoginProviders() []LoginProvider {
	return []LoginProvider{
		LoginProviderGoogle,
	}
}
