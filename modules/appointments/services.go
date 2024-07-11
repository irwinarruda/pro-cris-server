package appointments

import "github.com/irwinarruda/pro-cris-server/libs/proinject"

type AppointmentService struct {
	AppointmentRepository IAppointmentRepository `inject:"appointment_repository"`
}

func NewAppointmentService() *AppointmentService {
	return proinject.Resolve(&AppointmentService{})
}
