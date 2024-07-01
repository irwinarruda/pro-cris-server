package auth

type LoginTeacherDTO struct {
	Provider string `json:"provider" validate:"required,eq=google"`
	Token    string `json:"token" validate:"required,jwt"`
}

type CreateUserDTO struct {
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Picutre       *string `json:"picture"`
	EmailVerified bool    `json:"emailVerified"`
}
