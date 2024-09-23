package integration

import (
	"net/http"
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	authresources "github.com/irwinarruda/pro-cris-server/modules/auth/resources"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/providers"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"github.com/stretchr/testify/assert"
)

var validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.4S5J1zQJf4zZ2JZ9"
var invalidToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.4S5J1zQJf4zZ2JZ8"

func TestAuthServiceHappyPath(t *testing.T) {
	beforeEachAuth()

	var assert = assert.New(t)
	var authService = auth.NewAuthService()
	account1, _ := authService.Login(auth.LoginDTO{
		Provider: auth.LoginProviderGoogle,
		Token:    validToken,
	})

	account, err := authService.GetAccountByID(account1.ID)
	assert.NoError(err, "Should return the account.")
	assert.NotEqual(0, account1, "Should return a valid account id.")
	assert.Equal(account1.ID, account.ID, "Should return the account id.")
	assert.Equal("John Doe", account.Name, "Should return the correct account name.")
	assert.Equal("john@doe.com", account.Email, "Should return the correct account email.")
	assert.Equal(utils.StringP("https://www.google.com"), account.Picture, "Should return the correct account picture.")
	assert.Equal(false, account.EmailVerified, "Should return the correct account email verification status.")
	assert.Equal(auth.LoginProviderGoogle, account.Provider, "Should return the correct account provider.")
	assert.Equal(false, account.IsDeleted, "Should return the correct account deletion status.")

	account2, err := authService.EnsureAuthenticated(validToken, auth.LoginProviderGoogle)
	assert.NoError(err, "Should not return error with valid token")
	assert.Equal(account.ID, account2, "Should return the same ID as the created account.")

	afterEachAuth()
}

func TestAuthServiceErrorPath(t *testing.T) {
	beforeEachAuth()

	var assert = assert.New(t)
	var authService = auth.NewAuthService()

	_, err := authService.Login(auth.LoginDTO{
		Provider: auth.LoginProviderGoogle,
		Token:    "invalid",
	})
	assert.Error(err, "Should return an error with invalid access token.")

	_, err = authService.Login(auth.LoginDTO{
		Provider: auth.LoginProviderGoogle,
		Token:    invalidToken,
	})
	assert.Error(err, "Should return an error with wrong access token.")

	_, err = authService.Login(auth.LoginDTO{
		Provider: "invalid_provider",
		Token:    invalidToken,
	})
	assert.Error(err, "Should return an error with invalid provider.")

	u1, err := authService.Login(auth.LoginDTO{
		Provider: auth.LoginProviderGoogle,
		Token:    validToken,
	})
	assert.NoError(err, "Should not return error when login with new Account.")
	u2, err := authService.Login(auth.LoginDTO{
		Provider: auth.LoginProviderGoogle,
		Token:    validToken,
	})
	assert.NoError(err, "Should not return error when login with existing Account.")
	assert.Equal(u1.ID, u2.ID, "Should return same Account id when multiple logins.")

	afterEachAuth()
}

type MockGoogle struct{}

func (m *MockGoogle) Validate(token string) (providers.IGoogleUser, error) {
	if token == validToken {
		return providers.IGoogleUser{
			Email:         "john@doe.com",
			Name:          "John Doe",
			Picture:       "https://www.google.com",
			EmailVerified: false,
		}, nil
	}
	return providers.IGoogleUser{}, utils.NewAppError("Invalid google access token.", true, http.StatusBadRequest)
}

func beforeEachAuth() {
	proinject.Register("validate", configs.GetValidate(
		auth.GetLoginProviders(),
		[]string{},
		[]string{},
		[]string{},
	))
	proinject.Register("env", configs.GetEnv("../../.env"))
	proinject.Register("db", configs.GetDb())
	proinject.Register("google", &MockGoogle{})
	var authRepository = authresources.NewDbAuthRepository()
	proinject.Register("auth_repository", authRepository)
	authRepository.ResetAuth()
}

func afterEachAuth() {
	var authRepository = authresources.NewDbAuthRepository()
	authRepository.ResetAuth()
}
