package appointments

type IAppointmentRepository interface {
	CreateAppointmentsByRoutine() (int, error)
	CreateAppointment(appointment CreateAppointmentDTO) (int, error)
	GetAppointmentByID(id int) (Appointment, error)
	ResetAppointments()
}
