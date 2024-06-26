package auth

import (
	"fmt"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AuthRepository struct {
	Db configs.Db `inject:"db"`
}

func NewAuthRepository() *AuthRepository {
	return proinject.Resolve(&AuthRepository{})
}

func (a *AuthRepository) CreateUser(user CreateUserDTO) (User, error) {
	sql := fmt.Sprintf(`
    INSERT INTO "user"(
      name,
      email,
      picture,
      email_verified,
      provider
    ) %s
    RETURNING *;
	`, utils.SqlValues(1, 5))
	userE := UserEntity{}
	err := a.Db.Raw(sql, user.Name, user.Email, user.Picture, user.EmailVerified, user.Provider).Scan(&userE).Error
	if err != nil {
		return User{}, utils.NewAppError("Unable to create new user.", false, err)
	}
	return userE.ToUser(), nil
}

func (a *AuthRepository) GetUserByID(id int) (User, error) {
	usersE := []UserEntity{}
	a.Db.Raw("SELECT * FROM \"user\" WHERE id = ?;", id).Scan(&usersE)
	if len(usersE) == 0 {
		return User{}, utils.NewAppError("User not found.", true, nil)
	}
	return usersE[0].ToUser(), nil
}

func (a *AuthRepository) GetUserByEmail(email string) (User, error) {
	usersE := []UserEntity{}
	a.Db.Raw("SELECT * FROM \"user\" WHERE email = ?;", email).Scan(&usersE)
	if len(usersE) == 0 {
		return User{}, utils.NewAppError("User not found.", true, nil)
	}
	return usersE[0].ToUser(), nil
}

func (a *AuthRepository) GetIDByEmail(email string) (int, error) {
	ids := []int{}
	a.Db.Raw("SELECT id FROM \"user\" WHERE email = ?;", email).Scan(&ids)
	if len(ids) == 0 {
		return 0, utils.NewAppError("User not found.", true, nil)
	}
	return ids[0], nil
}

func (a *AuthRepository) ResetAuth() {
	a.Db.Exec("DELETE FROM \"user\";")
}
