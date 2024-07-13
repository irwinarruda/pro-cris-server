package appointments

import "github.com/irwinarruda/pro-cris-server/libs/proinject"

type AppointmentService struct {
	AppointmentRepository IAppointmentRepository `inject:"appointment_repository"`
}

func NewAppointmentService() *AppointmentService {
	return proinject.Resolve(&AppointmentService{})
}

func (a *AppointmentService) GetAppointmentByID(id int) (Appointment, error) {
	return a.AppointmentRepository.GetAppointmentByID(id)
}

func (a *AppointmentService) CreateAppointment(appointment CreateAppointmentDTO) (int, error) {
	return a.AppointmentRepository.CreateAppointment(appointment), nil
}
