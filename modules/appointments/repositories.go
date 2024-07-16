package appointments

type IAppointmentRepository interface {
	GetAppointmentByID(id int) (Appointment, error)
	CreateAppointmentsByRoutine() (int, error)
	CreateAppointment(appointment CreateAppointmentDTO) (int, error)
	UpdateAppointment(appointment UpdateAppointmentDTO) (int, error)
	ResetAppointments()
}
