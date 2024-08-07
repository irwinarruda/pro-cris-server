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

func (a *AuthService) Login(credentials LoginDTO) (Account, error) {
	if credentials.Provider != LoginProviderGoogle {
		return Account{}, utils.NewAppError("Invalid provider.", true, nil)
	}
	googleAccount, err := a.GoogleClient.Validate(credentials.Token)
	if err != nil {
		return Account{}, utils.NewAppError("Invalid google access token.", true, err)
	}
	existingAccount, err := a.AuthRepository.GetAccountByEmail(googleAccount.Email)
	if err == nil {
		return existingAccount, nil
	}
	account, err := a.AuthRepository.CreateAccount(CreateAccountDTO{
		Email:         googleAccount.Email,
		Name:          googleAccount.Name,
		Picture:       &googleAccount.Picture,
		EmailVerified: googleAccount.EmailVerified,
		Provider:      credentials.Provider,
	})
	return account, err
}

func (a *AuthService) EnsureAuthenticated(token string, provider LoginProvider) (int, error) {
	if provider != LoginProviderGoogle {
		return 0, utils.NewAppError("Invalid provider.", true, nil)
	}
	googleAccount, err := a.GoogleClient.Validate(token)
	if err != nil {
		return 0, utils.NewAppError(err.Error(), false, err)
	}
	return a.AuthRepository.GetIDByEmail(googleAccount.Email)
}

func (a *AuthService) GetAccountByID(id int) (Account, error) {
	return a.AuthRepository.GetAccountByID(id)
}
