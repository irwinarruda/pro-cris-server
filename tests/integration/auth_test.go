package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/providers"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"github.com/stretchr/testify/assert"
)

func TestAuthServiceHappyPath(t *testing.T) {
	beforeEachAuth()

	var assert = assert.New(t)
	var authService = auth.NewAuthService()
	user1, _ := authService.Login(auth.LoginDTO{
		Provider: auth.Google,
		Token:    "valid",
	})

	user, err := authService.GetUserByID(user1.ID)
	assert.NoError(err, "Should return the user.")
	assert.NotEqual(0, user1, "Should return a valid user id.")
	assert.Equal(user1.ID, user.ID, "Should return the user id.")
	assert.Equal("John Doe", user.Name, "Should return the correct user name.")
	assert.Equal("john@doe.com", user.Email, "Should return the correct user email.")
	assert.Equal(utils.StringP("https://www.google.com"), user.Picture, "Should return the correct user picture.")
	assert.Equal(false, user.EmailVerified, "Should return the correct user email verification status.")
	assert.Equal(auth.Google, user.Provider, "Should return the correct user provider.")
	assert.Equal(false, user.IsDeleted, "Should return the correct user deletion status.")

	user2, err := authService.EnsureAuthenticated("valid", auth.Google)
	assert.NoError(err, "Should not return error with valid token")
	assert.Equal(user.ID, user2, "Should return the same ID as the created user.")

	afterEachAuth()
}

func TestAuthServiceErrorPath(t *testing.T) {
	beforeEachAuth()

	var assert = assert.New(t)
	var authService = auth.NewAuthService()
	_, err := authService.Login(auth.LoginDTO{
		Provider: auth.Google,
		Token:    "invalid",
	})
	assert.Error(err, "Should return an error with invalid access token.")

	_, err = authService.Login(auth.LoginDTO{
		Provider: "invalid_provider",
		Token:    "invalid",
	})
	assert.Error(err, "Should return an error with invalid provider.")

	u1, err := authService.Login(auth.LoginDTO{
		Provider: auth.Google,
		Token:    "valid",
	})
	assert.NoError(err, "Should not return error when login with new User.")
	u2, err := authService.Login(auth.LoginDTO{
		Provider: auth.Google,
		Token:    "valid",
	})
	assert.NoError(err, "Should not return error when login with existing User.")
	assert.Equal(u1.ID, u2.ID, "Should return same User id when multiple logins.")

	afterEachAuth()
}

type MockGoogle struct{}

func (m *MockGoogle) Validate(token string) (providers.IGoogleUser, error) {
	if token == "valid" {
		return providers.IGoogleUser{
			Email:         "john@doe.com",
			Name:          "John Doe",
			Picture:       "https://www.google.com",
			EmailVerified: false,
		}, nil
	}
	return providers.IGoogleUser{}, utils.NewAppError("Invalid google access token.", true, nil)
}

func beforeEachAuth() {
	proinject.Register("env", configs.GetEnv("../../.env"))
	proinject.Register("db", configs.GetDb())
	proinject.Register("google", &MockGoogle{})
	var authRepository = auth.NewAuthRepository()
	proinject.Register("auth_repository", authRepository)
	authRepository.ResetAuth()
}

func afterEachAuth() {
	var authRepository = auth.NewAuthRepository()
	authRepository.ResetAuth()
}
