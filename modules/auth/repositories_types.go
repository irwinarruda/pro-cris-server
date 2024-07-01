package auth

type IAuthRepository interface {
	CreateUser(user CreateUserDTO) (int, error)
}
