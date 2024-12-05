package containers

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	appointmentsresources "github.com/irwinarruda/pro-cris-server/modules/appointments/resources"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	authresources "github.com/irwinarruda/pro-cris-server/modules/auth/resources"
	"github.com/irwinarruda/pro-cris-server/modules/settlements"
	settlementsresources "github.com/irwinarruda/pro-cris-server/modules/settlements/resources"
	statusresources "github.com/irwinarruda/pro-cris-server/modules/status/resources"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	studentsresources "github.com/irwinarruda/pro-cris-server/modules/students/resources"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/providers"
)

func InitInjections() {
	proinject.Register("validate", configs.GetValidate(map[string][]string{
		"login_provider": auth.GetLoginProvidersString(),
		// Keep those enums here for now because they are quite important to the main entities.
		"payment_style":    models.GetPaymentStylesString(),
		"payment_type":     models.GetPaymentTypesString(),
		"settlement_style": models.GetSettlementStylesString(),
	}))
	proinject.Register("env", configs.GetEnv())
	proinject.Register("db", configs.GetDb())
	proinject.Register("google", providers.NewGoogleClient())
	proinject.Register("student_repository", studentsresources.NewDbStudentRepository())
	proinject.Register("student_service", students.NewStudentService())
	proinject.Register("appointment_repository", appointmentsresources.NewDbAppointmentRepository())
	proinject.Register("appointment_service", appointments.NewAppointmentService())
	proinject.Register("settlement_repository", settlementsresources.NewDbSettlementRepository())
	proinject.Register("settlement_service", settlements.NewSettlementService())
	proinject.Register("auth_repository", authresources.NewDbAuthRepository())
	proinject.Register("status_repository", statusresources.NewDbStatusRepository())
}
