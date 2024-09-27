package auth

import (
	"net/http"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/providers"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AuthService struct {
	AuthRepository IAuthRepository   `inject:"auth_repository"`
	GoogleClient   providers.IGoogle `inject:"google"`
	Validate       configs.Validate  `inject:"validate"`
}

type IAuthService = *AuthService

func NewAuthService() *AuthService {
	return proinject.Resolve(&AuthService{})
}

func (a *AuthService) Login(credentials LoginDTO) (Account, error) {
	if err := a.Validate.Struct(credentials); err != nil {
		return Account{}, err
	}
	if credentials.Provider != LoginProviderGoogle {
		return Account{}, utils.NewAppError("Invalid provider.", true, http.StatusBadRequest)
	}
	googleAccount, err := a.GoogleClient.Validate(credentials.Token)
	if err != nil {
		return Account{}, utils.NewAppError("Invalid google access token.", true, http.StatusBadRequest)
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
		return 0, utils.NewAppError("Invalid provider.", true, http.StatusBadRequest)
	}
	googleAccount, err := a.GoogleClient.Validate(token)
	if err != nil {
		return 0, utils.NewAppError(err.Error(), false, http.StatusBadRequest)
	}
	return a.AuthRepository.GetIDByEmail(googleAccount.Email)
}

func (a *AuthService) GetAccountByID(id int) (Account, error) {
	return a.AuthRepository.GetAccountByID(id)
}
