package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	"github.com/irwinarruda/pro-cris-server/modules/appointments/resources"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/modules/auth/resources"
	"github.com/irwinarruda/pro-cris-server/modules/calendar"
	"github.com/irwinarruda/pro-cris-server/modules/calendar/resources"
	"github.com/irwinarruda/pro-cris-server/modules/students/resources"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"github.com/stretchr/testify/assert"
)

func TestAppointmentServiceHappyPath(t *testing.T) {
	idStudent := beforeEachAppointment()

	var assert = assert.New(t)
	var appointmentService = appointments.NewAppointmentService()

	id1, _ := appointmentService.CreateAppointment(appointments.CreateAppointmentDTO{
		IDStudent: idStudent,
		Price:     200,
		Duration:  int(1.8e+6),
		StartHour: "13:00",
		IsExtra:   true,
		IsPaid:    false,
		CalendarDay: appointments.CreateAppointmentCalendarDayDTO{
			Day:   1,
			Month: 1,
			Year:  2024,
		},
	})
	appointment1, err := appointmentService.GetAppointmentByID(id1)
	assert.NoError(err, "Should return get the created appointment.")
	assert.Equal(200.0, appointment1.Price, "Should return Price.")
	assert.Equal(int(1.8e+6), appointment1.Duration, "Should return Duration.")
	assert.Equal(idStudent, appointment1.Student.ID, "Should return IDStudent.")
	assert.Equal("13:00", appointment1.StartHour, "Should return StartHour.")
	assert.Equal(true, appointment1.IsExtra, "Should return IsExtra.")
	assert.Equal(false, appointment1.IsPaid, "Should return IsPaid.")
	assert.Equal(1, appointment1.CalendarDay.Day, "Should return Day.")
	assert.Equal(1, appointment1.CalendarDay.Month, "Should return Month.")
	assert.Equal(2024, appointment1.CalendarDay.Year, "Should return Year.")
	assert.Equal("John Doe", appointment1.Student.Name, "Should return Student Name.")
	assert.Equal("#000000", appointment1.Student.DisplayColor, "Should return Student Display Color.")
	assert.Equal(utils.StringP("http://example.com/picture.jpg"), appointment1.Student.Picture, "Should return Student Picture.")

	id2, _ := appointmentService.UpdateAppointment(appointments.UpdateAppointmentDTO{
		ID:      id1,
		IsExtra: false,
		IsPaid:  true,
		Price:   300,
	})
	appointment2, err := appointmentService.GetAppointmentByID(id2)
	assert.NoError(err, "Should return get the updated appointment.")
	assert.Equal(300.0, appointment2.Price, "Should return Price.")
	assert.Equal(false, appointment2.IsExtra, "Should return IsExtra.")
	assert.Equal(true, appointment2.IsPaid, "Should return IsPaid.")
}

func TestAppointmentServiceErrorPath(t *testing.T) {

}

func beforeEachAppointment() int {
	proinject.Register("env", configs.GetEnv("../../.env"))
	proinject.Register("db", configs.GetDb())

	var appointmentRepository = appointmentsresources.NewDbAppointmentRepository()
	appointmentRepository.ResetAppointments()
	proinject.Register("appointment_repository", appointmentRepository)

	calendarRepository := calendarresources.NewDbCalendarRepository()
	calendarRepository.ResetCalendarDays()
	proinject.Register("calendar_repository", calendarRepository)
	proinject.Register("calendar_service", calendar.NewCalendarService())

	var authRepository = authresources.NewDbAuthRepository()
	authRepository.ResetAuth()
	account, _ := authRepository.CreateAccount(auth.CreateAccountDTO{
		Email:         "john@doe.com",
		Name:          "John Doe",
		Picture:       utils.StringP("https://www.google.com"),
		EmailVerified: false,
		Provider:      auth.LoginProviderGoogle,
	})

	var studentRepository = studentsresources.NewDbStudentRepository()
	studentRepository.ResetStudents()
	idStudent := studentRepository.CreateStudent(mockCreateStudentDTO(account.ID))

	return idStudent
}
