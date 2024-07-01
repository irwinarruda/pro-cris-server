package auth

import (
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/providers"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AuthService struct {
	AuthRepository IAuthRepository   `inject:"auth_repository"`
	GoogleClient   providers.IGoogle `inject:"google"`
}

func NewAuthService() *AuthService {
	return configs.ResolveInject(&AuthService{})
}

func (a *AuthService) Login(credentials LoginTeacherDTO) (int, error) {
	if credentials.Provider != "google" {
		return 0, utils.NewAppError("Invalid provider", true, nil)
	}
	_, err := a.GoogleClient.Validate(credentials.Token)
	if err != nil {
		return 0, utils.NewAppError("Invalid google access token.", true, err)
	}

	return 0, nil
}
