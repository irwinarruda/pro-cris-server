package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/modules/settlements"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"github.com/stretchr/testify/assert"
)

func TestSettlementService(t *testing.T) {
	Init()
	var assert = assert.New(t)
	var settlementService = settlements.NewSettlementService()

	t.Run("Happy Path", func(t *testing.T) {
		idAccount, idStudent := beforeEachSettlement()
		settlementService.CreateSettlement(settlements.CreateSettlementDTO{
			IDAccount: idAccount,
			IDStudent: idStudent,
		})
		assert.Equal(1+1, 2, "1 + 1 should be equal to 2")
		afterEachSettlement()
	})

	t.Run("Error Path", func(t *testing.T) {
		beforeEachSettlement()
		afterEachSettlement()
	})

}

func beforeEachSettlement() (idAccount int, idStudent int) {
	var authRepository = proinject.Get[auth.IAuthRepository]("auth_repository")
	var studentRepository = proinject.Get[students.IStudentRepository]("student_repository")
	var settlementRepository = proinject.Get[settlements.ISettlementRepository]("settlement_repository")
	settlementRepository.ResetSettlement()
	authRepository.ResetAuth()
	studentRepository.ResetStudents()

	account, _ := authRepository.CreateAccount(auth.CreateAccountDTO{
		Email:         "john@doe.com",
		Name:          "John Doe",
		Picture:       utils.ToP("https://www.google.com"),
		EmailVerified: false,
		Provider:      auth.LoginProviderGoogle,
	})

	idStudent = studentRepository.CreateStudent(mockCreateStudentDTO(account.ID))
	return account.ID, idStudent
}

func afterEachSettlement() {
	var settlementRepository = proinject.Get[settlements.ISettlementRepository]("settlement_repository")
	settlementRepository.ResetSettlement()
}
