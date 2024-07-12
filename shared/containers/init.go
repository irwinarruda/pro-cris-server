package containers

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/modules/status"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/modules/students/resources"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/providers"
)

func InitInjections() {
	proinject.Register("env", configs.GetEnv())
	proinject.Register("validate", configs.GetValidate(
		auth.GetLoginProviders(),
		students.GetPaymentStyles(),
		students.GetPaymentTypes(),
		students.GetSettlementStyles(),
	))
	proinject.Register("db", configs.GetDb())
	proinject.Register("google", providers.NewGoogleClient())
	proinject.Register("students_repository", studentsresources.NewDbStudentRepository())
	proinject.Register("status_repository", status.NewDbStatusRepository())
	proinject.Register("appointment_repository", appointments.NewDbAppointmentRepository())
	proinject.Register("auth_repository", auth.NewDbAuthRepository())
}
