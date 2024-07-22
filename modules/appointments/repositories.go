package appointments

type IAppointmentRepository interface {
	GetAppointmentByID(data GetAppointmentDTO) (Appointment, error)
	CreateAppointmentsByRoutine() (int, error)
	CreateAppointment(appointment CreateAppointmentDTO) (int, error)
	UpdateAppointment(appointment UpdateAppointmentDTO) (int, error)
	DeleteAppointment(data DeleteAppointmentDTO) (int, error)
	ResetAppointments()
}
