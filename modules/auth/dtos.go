package auth

type LoginDTO struct {
	Provider LoginProvider `json:"provider" validate:"required,login_provider"`
	Token    string        `json:"token" validate:"required,jwt"`
}

type CreateAccountDTO struct {
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	Picture       *string       `json:"picture"`
	EmailVerified bool          `json:"emailVerified"`
	Provider      LoginProvider `json:"provider"`
}
