package appointments

type IAppointmentRepository interface {
	CreateAppointment(appointment CreateAppointmentDTO) int
	GetAppointmentByID(id int) (Appointment, error)
}
