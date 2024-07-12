package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	"github.com/irwinarruda/pro-cris-server/modules/appointments/resources"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/modules/auth/resources"
	"github.com/irwinarruda/pro-cris-server/modules/students/resources"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"github.com/stretchr/testify/assert"
)

func TestAppointmentServiceHappyPath(t *testing.T) {
	beforeEachAppointment()

	var assert = assert.New(t)
	var _ = appointments.NewAppointmentService()
	assert.Equal(0, 0)
}

func TestAppointmentServiceErrorPath(t *testing.T) {
	beforeEachAppointment()
}

func beforeEachAppointment() {
	proinject.Register("env", configs.GetEnv("../../.env"))
	proinject.Register("db", configs.GetDb())
	proinject.Register("appointment_repository", appointmentsresources.NewDbAppointmentRepository())
	var studentRepository = studentsresources.NewDbStudentRepository()
	studentRepository.ResetStudents()
	var authRepository = authresources.NewDbAuthRepository()
	authRepository.ResetAuth()
	authRepository.CreateAccount(auth.CreateAccountDTO{
		Email:         "john@doe.com",
		Name:          "John Doe",
		Picture:       utils.StringP("https://www.google.com"),
		EmailVerified: false,
		Provider:      auth.LoginProviderGoogle,
	})
}
