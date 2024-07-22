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

func (a *AppointmentService) GetAppointmentByID(data GetAppointmentDTO) (Appointment, error) {
	return a.AppointmentRepository.GetAppointmentByID(data)
}

func (a *AppointmentService) UpdateAppointment(appointment UpdateAppointmentDTO) (int, error) {
	return a.AppointmentRepository.UpdateAppointment(appointment)
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

func (a *AppointmentService) DeleteAppointment(data DeleteAppointmentDTO) (int, error) {
	return a.AppointmentRepository.DeleteAppointment(data)
}
