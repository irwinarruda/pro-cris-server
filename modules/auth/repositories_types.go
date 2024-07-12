package auth

import "time"

type IAuthRepository interface {
	CreateUser(user CreateUserDTO) (User, error)
	GetUserByID(id int) (User, error)
	GetUserByEmail(email string) (User, error)
	GetIDByEmail(email string) (int, error)
	ResetAuth()
}

type DbUser struct {
	ID            int
	Name          string
	Email         string
	EmailVerified bool
	Picture       *string
	Provider      LoginProvider
	IsDeleted     bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (u *DbUser) ToUser() User {
	return User{
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
