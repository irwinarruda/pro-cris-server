package appointments

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

type AppointmentRepository struct {
	Db configs.Db `inject:"db"`
}

func NewAppointmentRepository() *AppointmentRepository {
	return proinject.Resolve(&AppointmentRepository{})
}

func (a *AppointmentRepository) CreateAppointment(appointment CreateAppointmentDTO) int {
	return 0
}

func (a *AppointmentRepository) GetAppointmentByID(id int) (Appointment, error) {
	return Appointment{}, nil
}
