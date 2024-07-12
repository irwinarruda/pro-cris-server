package appointmentsresources

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

type DbAppointmentRepository struct {
	Db configs.Db `inject:"db"`
}

func NewDbAppointmentRepository() *DbAppointmentRepository {
	return proinject.Resolve(&DbAppointmentRepository{})
}

func (a *DbAppointmentRepository) CreateAppointment(appointment appointments.CreateAppointmentDTO) int {
	return 0
}

func (a *DbAppointmentRepository) GetAppointmentByID(id int) (appointments.Appointment, error) {
	return appointments.Appointment{}, nil
}
