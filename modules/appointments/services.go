package appointments

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/calendar"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AppointmentService struct {
	AppointmentRepository IAppointmentRepository    `inject:"appointment_repository"`
	CalendarService       *calendar.CalendarService `inject:"calendar_service"`
}

func NewAppointmentService() *AppointmentService {
	return proinject.Resolve(&AppointmentService{})
}

func (a *AppointmentService) GetAppointmentByID(id int) (Appointment, error) {
	appointment, err := a.AppointmentRepository.GetAppointmentByID(id)
	if err != nil {
		return Appointment{}, utils.NewAppError("Appointment not found.", true, err)
	}
	return appointment, nil
}

func (a *AppointmentService) CreateAppointment(appointment CreateAppointmentDTO) (int, error) {
	id, err := a.CalendarService.CreateCalendarDayIfNotExists(appointment.CalendarDay.Day, appointment.CalendarDay.Month, appointment.CalendarDay.Year)
	if err != nil {
		return 0, err
	}
	appointment.CalendarDay.ID = id
	id, err = a.AppointmentRepository.CreateAppointment(appointment)
	if err != nil {
		return 0, utils.NewAppError("Error creating appointment.", false, nil)
	}
	return id, nil
}
