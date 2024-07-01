package auth

type LoginTeacherDTO struct {
	Provider LoginProvider `json:"provider" validate:"required,eq=Google"`
	Token    string        `json:"token" validate:"required,jwt"`
}

type CreateUserDTO struct {
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	Picture       *string       `json:"picture"`
	EmailVerified bool          `json:"emailVerified"`
	Provider      LoginProvider `json:"provider"`
}
