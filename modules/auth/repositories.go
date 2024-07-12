package auth

type IAuthRepository interface {
	CreateAccount(account CreateAccountDTO) (Account, error)
	GetAccountByID(id int) (Account, error)
	GetAccountByEmail(email string) (Account, error)
	GetIDByEmail(email string) (int, error)
	ResetAuth()
}
