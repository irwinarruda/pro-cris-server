package containers

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/modules/status"
	"github.com/irwinarruda/pro-cris-server/modules/students"
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
	proinject.Register("students_repository", students.NewStudentRepository())
	proinject.Register("status_repository", status.NewStatusRepository())
	proinject.Register("appointment_repository", appointments.NewAppointmentRepository())
	proinject.Register("auth_repository", auth.NewAuthRepository())
}
