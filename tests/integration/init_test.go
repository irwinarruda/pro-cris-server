package integration

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	appointmentsresources "github.com/irwinarruda/pro-cris-server/modules/appointments/resources"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	authresources "github.com/irwinarruda/pro-cris-server/modules/auth/resources"
	statusresources "github.com/irwinarruda/pro-cris-server/modules/status/resources"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	studentsresources "github.com/irwinarruda/pro-cris-server/modules/students/resources"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

func Init() {
	proinject.Register("validate", configs.GetValidate(map[string][]string{
		"login_provider":   auth.GetLoginProvidersString(),
		"payment_style":    students.GetPaymentStylesString(),
		"payment_type":     students.GetPaymentTypesString(),
		"settlement_style": students.GetSettlementStylesString(),
	}))
	proinject.Register("env", configs.GetEnv("../../.env"))
	proinject.Register("db", configs.GetDb())
	proinject.Register("google", &MockGoogle{})
	proinject.Register("auth_repository", authresources.NewDbAuthRepository())
	proinject.Register("appointment_repository", appointmentsresources.NewDbAppointmentRepository())
	proinject.Register("auth_repository", authresources.NewDbAuthRepository())
	proinject.Register("status_repository", statusresources.NewDbStatusRepository())
	proinject.Register("student_repository", studentsresources.NewDbStudentRepository())
	proinject.Register("students_service", students.NewStudentService())
}
