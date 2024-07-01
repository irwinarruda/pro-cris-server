package auth

import "github.com/irwinarruda/pro-cris-server/shared/configs"

type AuthRepository struct {
	Db configs.Db `inject:"db"`
}

func NewAuthRepository() *AuthRepository {
	return configs.ResolveInject(&AuthRepository{})
}

func CreateUser(user CreateUserDTO) (int, error) {
	return 0, nil
}
