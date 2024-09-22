package containers

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	appointmentsresources "github.com/irwinarruda/pro-cris-server/modules/appointments/resources"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	authresources "github.com/irwinarruda/pro-cris-server/modules/auth/resources"
	statusresources "github.com/irwinarruda/pro-cris-server/modules/status/resources"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	studentsresources "github.com/irwinarruda/pro-cris-server/modules/students/resources"
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
	proinject.Register("appointment_repository", appointmentsresources.NewDbAppointmentRepository())
	proinject.Register("auth_repository", authresources.NewDbAuthRepository())
	proinject.Register("status_repository", statusresources.NewDbStatusRepository())
	proinject.Register("students_repository", studentsresources.NewDbStudentRepository())
	proinject.Register("students_service", students.NewStudentService())
}
