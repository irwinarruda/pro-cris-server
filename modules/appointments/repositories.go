package appointments

type IAppointmentRepository interface {
	CreateAppointmentsByRoutine() int
	CreateAppointment(appointment CreateAppointmentDTO) int
	GetAppointmentByID(id int) (Appointment, error)
	ResetAppointments()
}
