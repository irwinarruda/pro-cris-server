package integration

import (
	"testing"
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"github.com/stretchr/testify/assert"
)

func TestAppointmentService(t *testing.T) {
	Init()
	t.Run("Happy Path", func(t *testing.T) {
		idAccount, idStudent, idStudent2 := beforeEachAppointment()

		var assert = assert.New(t)
		var appointmentService = appointments.NewAppointmentService()

		id1, _ := appointmentService.CreateAppointment(mockCreateAppointmentDTO(idAccount, idStudent))
		appointment1, err := appointmentService.GetAppointmentByID(appointments.GetAppointmentDTO{
			IDAccount: idAccount,
			ID:        id1,
		})
		assert.NoError(err, "Should return get the created appointment.")
		assert.Equal(200.0, appointment1.Price, "Should return Price.")
		assert.Equal(int(1.8e+6), appointment1.Duration, "Should return Duration.")
		assert.Equal(idStudent, appointment1.Student.ID, "Should return IDStudent.")
		assert.Equal(13, appointment1.StartHour, "Should return StartHour.")
		assert.Equal(true, appointment1.IsExtra, "Should return IsExtra.")
		assert.Equal(false, appointment1.IsPaid, "Should return IsPaid.")
		assert.Equal(1, appointment1.CalendarDay.Day(), "Should return Day.")
		assert.Equal(time.Month(1), appointment1.CalendarDay.Month(), "Should return Month.")
		assert.Equal(2024, appointment1.CalendarDay.Year(), "Should return Year.")
		assert.Equal("John Doe", appointment1.Student.Name, "Should return Student Name.")
		assert.Equal("#000000", appointment1.Student.DisplayColor, "Should return Student Display Color.")
		assert.Equal(utils.ToP("http://example.com/picture.jpg"), appointment1.Student.Picture, "Should return Student Picture.")

		id2, _ := appointmentService.UpdateAppointment(appointments.UpdateAppointmentDTO{
			IDAccount: idAccount,
			ID:        id1,
			IsExtra:   false,
			IsPaid:    true,
			Price:     300,
		})
		appointment2, err := appointmentService.GetAppointmentByID(appointments.GetAppointmentDTO{
			IDAccount: idAccount,
			ID:        id2,
		})
		assert.NotEqual(appointment2.CreatedAt, appointment2.UpdatedAt, "UpdatedAt should be updated.")
		assert.NoError(err, "Should return get the updated appointment.")
		assert.Equal(300.0, appointment2.Price, "Should return Price.")
		assert.Equal(false, appointment2.IsExtra, "Should return IsExtra.")
		assert.Equal(true, appointment2.IsPaid, "Should return IsPaid.")

		id3, err := appointmentService.DeleteAppointment(appointments.DeleteAppointmentDTO{
			IDAccount: idAccount,
			ID:        id2,
		})
		assert.NoError(err, "Should not return error deleting appointment.")
		assert.Equal(id2, id3, "Should return the same id deleted.")
		_, err = appointmentService.GetAppointmentByID(appointments.GetAppointmentDTO{
			IDAccount: idAccount,
			ID:        id3,
		})
		assert.Error(err, "Should return error because the appointment was deleted.")

		date, err := time.Parse(time.DateOnly, "2024-01-02")
		createdAppointments, _ := appointmentService.CreateDailyAppointmentsByStudentsRoutine(appointments.CreateDailyAppointmentsByStudentsRoutineDTO{
			IDAccount:   idAccount,
			CalendarDay: date,
		})
		assert.NoError(err, "Should not return error creating appointments.")
		assert.Len(createdAppointments, 2, "Should create 2 appointments from the students.")
		assert.Equal(idStudent, createdAppointments[0].Student.ID, "Should return the ID from the student routine.")
		assert.Equal(100.0, createdAppointments[0].Price, "Should return the price from the student routine.")
		assert.Equal(60, createdAppointments[0].Duration, "Should return the duration from the student routine.")
		assert.Equal(8, createdAppointments[0].StartHour, "Should return the start hour from the student routine.")
		assert.Equal(idStudent2, createdAppointments[1].Student.ID, "Should return the ID from the student routine.")
		assert.Equal(300.0, createdAppointments[1].Price, "Should return the price from the student routine.")
		assert.Equal(120, createdAppointments[1].Duration, "Should return the duration from the student routine.")
		assert.Equal(10, createdAppointments[1].StartHour, "Should return the start hour from the student routine.")

		afterEachAppointment()
	})

	t.Run("Error Path", func(t *testing.T) {
		idAccount, idStudent, _ := beforeEachAppointment()

		var assert = assert.New(t)
		var appointmentService = appointments.NewAppointmentService()

		_, err := appointmentService.GetAppointmentByID(appointments.GetAppointmentDTO{
			IDAccount: idAccount,
			ID:        8,
		})
		assert.Error(err, "Should return error when appointment/account not found.")

		_, err = appointmentService.CreateAppointment(mockCreateAppointmentDTO(idAccount, 8))
		assert.Error(err, "Should return error when student/account  does not exist.")
		_, err = appointmentService.CreateAppointment(mockCreateAppointmentDTO(8, idStudent))
		assert.Error(err, "Should return error when student/account  does not exist.")

		idAppointment, _ := appointmentService.CreateAppointment(mockCreateAppointmentDTO(idAccount, idStudent))
		_, err = appointmentService.DeleteAppointment(appointments.DeleteAppointmentDTO{
			IDAccount: idAccount,
			ID:        8,
		})
		assert.Error(err, "Should return error when appointment/account does not exist.")
		_, err = appointmentService.DeleteAppointment(appointments.DeleteAppointmentDTO{
			IDAccount: 8,
			ID:        idAppointment,
		})
		assert.Error(err, "Should return error when appointment/account does not exist.")

		afterEachAppointment()
	})
}

func mockCreateAppointmentDTO(idAccount, idStudent int) appointments.CreateAppointmentDTO {
	date, _ := time.Parse(time.DateOnly, "2024-01-01")
	return appointments.CreateAppointmentDTO{
		IDAccount:   idAccount,
		IDStudent:   idStudent,
		Price:       200,
		Duration:    int(1.8e+6),
		StartHour:   13,
		IsExtra:     true,
		IsPaid:      false,
		CalendarDay: date,
	}
}

func beforeEachAppointment() (idAccount int, idStudent int, idStudent2 int) {
	var appointmentRepository = proinject.Get[appointments.IAppointmentRepository]("appointment_repository")
	var authRepository = proinject.Get[auth.IAuthRepository]("auth_repository")
	var studentRepository = proinject.Get[students.IStudentRepository]("student_repository")
	appointmentRepository.ResetAppointments()
	authRepository.ResetAuth()
	studentRepository.ResetStudents()

	account, _ := authRepository.CreateAccount(auth.CreateAccountDTO{
		Email:         "john@doe.com",
		Name:          "John Doe",
		Picture:       utils.ToP("https://www.google.com"),
		EmailVerified: false,
		Provider:      auth.LoginProviderGoogle,
	})

	idStudent = studentRepository.CreateStudent(mockCreateStudentDTO(account.ID))
	idStudent2 = studentRepository.CreateStudent(mockCreateStudentDTO2(account.ID))
	return account.ID, idStudent, idStudent2
}

func afterEachAppointment() {
	var authRepository = proinject.Get[auth.IAuthRepository]("auth_repository")
	var studentRepository = proinject.Get[students.IStudentRepository]("student_repository")
	var appointmentRepository = proinject.Get[appointments.IAppointmentRepository]("appointment_repository")
	authRepository.ResetAuth()
	studentRepository.ResetStudents()
	appointmentRepository.ResetAppointments()
}
