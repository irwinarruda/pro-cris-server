package auth

type IAuthRepository interface {
	CreateUser(user CreateUserDTO) (User, error)
	GetUserByID(id int) (User, error)
	GetUserByEmail(email string) (User, error)
	GetIDByEmail(email string) (int, error)
	ResetAuth()
}
