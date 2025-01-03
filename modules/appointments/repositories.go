package appointments

type IAppointmentRepository interface {
	GetAppointmentByID(data GetAppointmentDTO) (Appointment, error)
	GetAppointmentsByID(data GetAppointmentsDTO) ([]Appointment, error)
	GetNotSettledAppointmentsByStudent(data GetNotSettledAppointmentsByStudentDTO) ([]Appointment, error)
	GetAppointmentsByDateRange(data GetAppointmentsByDateRangeDTO) ([]Appointment, error)
	CreateAppointment(appointment CreateAppointmentDTO) (int, error)
	UpdateAppointment(appointment UpdateAppointmentDTO) (int, error)
	DeleteAppointment(data DeleteAppointmentDTO) (int, error)
	ResetAppointments()
}
