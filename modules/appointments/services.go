package appointments

import (
	"net/http"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AppointmentService struct {
	Validate              configs.Validate         `inject:"validate"`
	AppointmentRepository IAppointmentRepository   `inject:"appointment_repository"`
	StudentService        *students.StudentService `inject:"students_service"`
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

func (a *AppointmentService) CreateDailyAppointmentsByStudentsRoutine(data CreateDailyAppointmentsByStudentsRoutineDTO) error {
	if err := a.Validate.Struct(data); err != nil {
		return err
	}
	return nil
}

func (a *AppointmentService) DeleteAppointment(data DeleteAppointmentDTO) (int, error) {
	if err := a.Validate.Struct(data); err != nil {
		return 0, err
	}
	return a.AppointmentRepository.DeleteAppointment(data)
}
