package auth

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/providers"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AuthService struct {
	AuthRepository IAuthRepository   `inject:"auth_repository"`
	GoogleClient   providers.IGoogle `inject:"google"`
}

func NewAuthService() *AuthService {
	return proinject.Resolve(&AuthService{})
}

func (a *AuthService) Login(credentials LoginTeacherDTO) (User, error) {
	if credentials.Provider != Google {
		return User{}, utils.NewAppError("Invalid provider.", true, nil)
	}
	googleUser, err := a.GoogleClient.Validate(credentials.Token)
	if err != nil {
		return User{}, utils.NewAppError("Invalid google access token.", true, err)
	}
	existingUser, err := a.AuthRepository.GetUserByEmail(googleUser.Email)
	if err == nil {
		return existingUser, nil
	}
	user, err := a.AuthRepository.CreateUser(CreateUserDTO{
		Email:         googleUser.Email,
		Name:          googleUser.Name,
		Picture:       &googleUser.Picture,
		EmailVerified: googleUser.EmailVerified,
	})
	return user, err
}

func (a *AuthService) EnsureAuthenticated(token string, provider LoginProvider) (int, error) {
	if provider != Google {
		return 0, utils.NewAppError("Invalid provider.", true, nil)
	}
	googleUser, err := a.GoogleClient.Validate(token)
	if err != nil {
		return 0, utils.NewAppError(err.Error(), false, err)
	}
	return a.AuthRepository.GetIdByEmail(googleUser.Email)
}

func (a *AuthService) GetUserByID(id int) (User, error) {
	return a.AuthRepository.GetUserByID(id)
}
