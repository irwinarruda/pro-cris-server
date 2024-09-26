package appointments

import (
	"net/http"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/date"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AppointmentService struct {
	Validate              configs.Validate         `inject:"validate"`
	AppointmentRepository IAppointmentRepository   `inject:"appointment_repository"`
	StudentService        *students.StudentService `inject:"students_service"`
	DateService           *date.DateService        `inject:"date_service"`
}

func NewAppointmentService() *AppointmentService {
	return proinject.Resolve(&AppointmentService{})
}

func (a *AppointmentService) GetAppointmentByID(data GetAppointmentDTO) (Appointment, error) {
	if err := a.Validate.Struct(data); err != nil {
		return Appointment{}, err
	}
	return a.AppointmentRepository.GetAppointmentByID(data)
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
	id, err := a.AppointmentRepository.CreateAppointment(appointment)
	if err != nil {
		return 0, utils.NewAppError("Error creating appointment.", false, http.StatusBadRequest)
	}
	return id, nil
}

func (a *AppointmentService) CreateDailyAppointmentsByStudentsRoutine(data CreateDailyAppointmentsByStudentsRoutineDTO) ([]Appointment, error) {
	createdAppointments := []Appointment{}
	// errorList := []string{}

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
