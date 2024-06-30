package auth

import "github.com/irwinarruda/pro-cris-server/shared/configs"

type AuthService struct {
}

func NewAuthService() *AuthService {
	return configs.ResolveInject(&AuthService{})
}

func (a *AuthService) Login(credentials interface{}) int {
	return 0
}

func (a *AuthService) GetUserByID(id int) (User, error) {
	return User{}, nil
}
