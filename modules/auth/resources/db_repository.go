package authresources

import (
	"fmt"
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type DbAccount struct {
	ID            int
	Name          string
	Email         string
	EmailVerified bool
	Picture       *string
	Provider      auth.LoginProvider
	IsDeleted     bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (u *DbAccount) ToAccount() auth.Account {
	return auth.Account{
		ID:            u.ID,
		Name:          u.Name,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		Picture:       u.Picture,
		Provider:      u.Provider,
		IsDeleted:     u.IsDeleted,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}
}

type DbAuthRepository struct {
	Db configs.Db `inject:"db"`
}

func NewDbAuthRepository() *DbAuthRepository {
	return proinject.Resolve(&DbAuthRepository{})
}

func (a *DbAuthRepository) CreateAccount(account auth.CreateAccountDTO) (auth.Account, error) {
	sql := fmt.Sprintf(`
    INSERT INTO "account"(
      name,
      email,
      picture,
      email_verified,
      provider
    ) %s
    RETURNING *;
	`, utils.SqlValues(1, 5))
	accountE := DbAccount{}
	err := a.Db.Raw(sql, account.Name, account.Email, account.Picture, account.EmailVerified, account.Provider).Scan(&accountE).Error
	if err != nil {
		return auth.Account{}, utils.NewAppError("Unable to create new account.", false, err)
	}
	return accountE.ToAccount(), nil
}

func (a *DbAuthRepository) GetAccountByID(id int) (auth.Account, error) {
	accountsE := []DbAccount{}
	a.Db.Raw("SELECT * FROM \"account\" WHERE id = ?;", id).Scan(&accountsE)
	if len(accountsE) == 0 {
		return auth.Account{}, utils.NewAppError("Account not found.", true, nil)
	}
	return accountsE[0].ToAccount(), nil
}

func (a *DbAuthRepository) GetAccountByEmail(email string) (auth.Account, error) {
	accountsE := []DbAccount{}
	a.Db.Raw("SELECT * FROM \"account\" WHERE email = ?;", email).Scan(&accountsE)
	if len(accountsE) == 0 {
		return auth.Account{}, utils.NewAppError("Account not found.", true, nil)
	}
	return accountsE[0].ToAccount(), nil
}

func (a *DbAuthRepository) GetIDByEmail(email string) (int, error) {
	ids := []int{}
	a.Db.Raw("SELECT id FROM \"account\" WHERE email = ?;", email).Scan(&ids)
	if len(ids) == 0 {
		return 0, utils.NewAppError("Account not found.", true, nil)
	}
	return ids[0], nil
}

func (a *DbAuthRepository) ResetAuth() {
	a.Db.Exec("DELETE FROM \"account\";")
	a.Db.Exec("ALTER SEQUENCE account_id_seq RESTART WITH 1;")
}
