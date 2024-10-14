package appointments

import (
	"net/http"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/date"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/constants"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AppointmentService struct {
	Validate              configs.Validate         `inject:"validate"`
	AppointmentRepository IAppointmentRepository   `inject:"appointment_repository"`
	StudentService        students.IStudentService `inject:"students_service"`
	DateService           date.IDateService        `inject:"date_service"`
}

type IAppointmentService = *AppointmentService

func NewAppointmentService() *AppointmentService {
	return proinject.Resolve(&AppointmentService{})
}

func (a *AppointmentService) GetAppointmentByID(data GetAppointmentDTO) (Appointment, error) {
	if err := a.Validate.Struct(data); err != nil {
		return Appointment{}, err
	}
	return a.AppointmentRepository.GetAppointmentByID(data)
}

func (a *AppointmentService) GetAppointmentsByDateRange(data GetAppointmentsByDateRangeDTO) ([]Appointment, error) {
	if err := a.Validate.Struct(data); err != nil {
		return []Appointment{}, err
	}
	return a.AppointmentRepository.GetAppointmentsByDateRange(data)
}

func (a *AppointmentService) UpdateAppointment(appointment UpdateAppointmentDTO) (int, error) {
	if err := a.Validate.Struct(appointment); err != nil {
		return 0, err
	}
	return a.AppointmentRepository.UpdateAppointment(appointment)
}

func (a *AppointmentService) CreateAppointment(appointment CreateAppointmentDTO) (int, error) {
	if err := a.Validate.Struct(appointment); err != nil {
		return 0, err
	}
	hasStudent := a.StudentService.DoesStudentExists(students.DoesStudentExistsDTO{
		IDAccount: appointment.IDAccount,
		ID:        appointment.IDStudent,
	})
	if !hasStudent {
		return 0, utils.NewAppError("Student not found.", true, http.StatusBadRequest)
	}
	initialDate := appointment.CalendarDay.AddDate(0, 0, -1)
	finalDate := appointment.CalendarDay.AddDate(0, 0, 1)
	appointmentsRange, err := a.AppointmentRepository.GetAppointmentsByDateRange(GetAppointmentsByDateRangeDTO{
		IDAccount:   appointment.IDAccount,
		IDStudent:   appointment.IDStudent,
		InitialDate: initialDate,
		FinalDate:   finalDate,
	})
	if err != nil {
		return 0, err
	}
	for _, appointmentRange := range appointmentsRange {
		if appointment.CalendarDay == appointmentRange.CalendarDay && utils.IsOverlapping(appointment.StartHour, appointment.Duration, appointmentRange.StartHour, appointmentRange.Duration) {
			return 0, utils.NewAppError("Appointment is overlapping with another appointment.", true, http.StatusBadRequest)
		}
		if appointment.CalendarDay.Before(appointmentRange.CalendarDay) && utils.IsOverlapping(appointment.StartHour, appointment.Duration, appointmentRange.StartHour+constants.Hour24, appointmentRange.Duration) {
			return 0, utils.NewAppError("Appointment is overlapping with another appointment.", true, http.StatusBadRequest)
		}
		if appointment.CalendarDay.After(appointmentRange.CalendarDay) && utils.IsOverlapping(appointment.StartHour+constants.Hour24, appointment.Duration, appointmentRange.StartHour, appointmentRange.Duration) {
			return 0, utils.NewAppError("Appointment is overlapping with another appointment.", true, http.StatusBadRequest)
		}
	}
	id, err := a.AppointmentRepository.CreateAppointment(appointment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AppointmentService) CreateDailyAppointmentsByStudentsRoutine(data CreateDailyAppointmentsByStudentsRoutineDTO) ([]Appointment, error) {
	createdAppointments := []Appointment{}

	if err := a.Validate.Struct(data); err != nil {
		return createdAppointments, err
	}
	students, err := a.StudentService.GetAllStudents(students.GetAllStudentsDTO{IDAccount: data.IDAccount})
	if err != nil {
		return createdAppointments, err
	}

	weekDay := a.DateService.GetWeekDayFromDate(data.CalendarDay)
	for _, student := range students {
		for _, routinePlan := range student.Routine {
			if weekDay != routinePlan.WeekDay {
				continue
			}
			id, err := a.CreateAppointment(CreateAppointmentDTO{
				IDAccount:   data.IDAccount,
				IDStudent:   student.ID,
				CalendarDay: data.CalendarDay,
				StartHour:   routinePlan.StartHour,
				Duration:    routinePlan.Duration,
				Price:       routinePlan.Price,
				IsExtra:     false,
				IsPaid:      false,
			})
			if err != nil {
				continue
			}
			appointment, err := a.GetAppointmentByID(GetAppointmentDTO{IDAccount: data.IDAccount, ID: id})
			if err != nil {
				continue
			}
			createdAppointments = append(createdAppointments, appointment)
		}
	}
	return createdAppointments, nil
}

func (a *AppointmentService) DeleteAppointment(data DeleteAppointmentDTO) (int, error) {
	if err := a.Validate.Struct(data); err != nil {
		return 0, err
	}
	return a.AppointmentRepository.DeleteAppointment(data)
}
