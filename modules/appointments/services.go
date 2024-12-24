package appointments

import (
	"net/http"
	"slices"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/constants"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AppointmentService struct {
	Validate              configs.Validate         `inject:"validate"`
	AppointmentRepository IAppointmentRepository   `inject:"appointment_repository"`
	StudentService        students.IStudentService `inject:"student_service"`
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

func (a *AppointmentService) GetNotSettledAppointmentsByStudent(data GetNotSettledAppointmentsByStudentDTO) ([]Appointment, error) {
	if err := a.Validate.Struct(data); err != nil {
		return []Appointment{}, err
	}
	return a.AppointmentRepository.GetNotSettledAppointmentsByStudent(data)
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
		InitialDate: initialDate,
		FinalDate:   finalDate,
	})
	if err != nil {
		return 0, err
	}
	for _, appointmentRange := range appointmentsRange {
		if appointment.CalendarDay == appointmentRange.CalendarDay && utils.IsOverlappingInt(appointment.StartHour, appointment.Duration, appointmentRange.StartHour, appointmentRange.Duration) {
			return 0, utils.NewAppErrors("Appointment is overlapping with another appointment.", appointmentRange, true, http.StatusBadRequest)
		}
		if appointment.CalendarDay.Before(appointmentRange.CalendarDay) && utils.IsOverlappingInt(appointment.StartHour, appointment.Duration, appointmentRange.StartHour+constants.Hour24, appointmentRange.Duration) {
			return 0, utils.NewAppErrors("Appointment is overlapping with another appointment.", appointmentRange, true, http.StatusBadRequest)
		}
		if appointment.CalendarDay.After(appointmentRange.CalendarDay) && utils.IsOverlappingInt(appointment.StartHour+constants.Hour24, appointment.Duration, appointmentRange.StartHour, appointmentRange.Duration) {
			return 0, utils.NewAppErrors("Appointment is overlapping with another appointment.", appointmentRange, true, http.StatusBadRequest)
		}
	}
	id, err := a.AppointmentRepository.CreateAppointment(appointment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AppointmentService) CreateDailyAppointmentsByStudentsRoutine(data CreateDailyAppointmentsByStudentsRoutineDTO) ([]int, error) {
	if err := a.Validate.Struct(data); err != nil {
		return []int{}, err
	}
	students, err := a.StudentService.GetAllStudents(students.GetAllStudentsDTO{IDAccount: data.IDAccount})
	if err != nil {
		return []int{}, err
	}

	createdAppointments := []int{}
	notCreatedRoutine := []int{}
	weekDay := models.FromTime(data.CalendarDay)
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
				notCreatedRoutine = append(notCreatedRoutine, routinePlan.ID)
				continue
			}
			createdAppointments = append(createdAppointments, id)
		}
	}
	if len(notCreatedRoutine) > 0 {
		meta := struct {
			CreatedAppointments   []int `json:"createdAppointments"`
			NotCreatedRoutinePlan []int `json:"notCreatedRoutinePlan"`
		}{CreatedAppointments: createdAppointments, NotCreatedRoutinePlan: notCreatedRoutine}
		return []int{}, utils.NewAppErrors("Some appointments were not created.", meta, true, http.StatusPartialContent)
	}
	return createdAppointments, nil
}

func (a *AppointmentService) DeleteAppointment(data DeleteAppointmentDTO) (int, error) {
	if err := a.Validate.Struct(data); err != nil {
		return 0, err
	}
	return a.AppointmentRepository.DeleteAppointment(data)
}

func (a *AppointmentService) DoAppointmentsExist(data DoAppointmentsExistDTO) (bool, error) {
	if err := a.Validate.Struct(data); err != nil {
		return false, err
	}
	appointments, err := a.AppointmentRepository.GetAppointmentsByID(GetAppointmentsDTO(data))
	meta := struct {
		NotExistingAppointments []int `json:"notExistingAppointments"`
	}{}
	if err != nil {
		return false, utils.NewAppErrors(err.Error(), meta, false, http.StatusBadGateway)
	}
	if len(data.IDs) == len(appointments) {
		return true, nil
	}
	for _, id := range data.IDs {
		if !slices.ContainsFunc(appointments, func(app Appointment) bool { return app.ID == id }) {
			meta.NotExistingAppointments = append(meta.NotExistingAppointments, id)
		}
	}
	return false, utils.NewAppErrors("Some appointments do not exist.", meta, true, http.StatusPartialContent)
}
