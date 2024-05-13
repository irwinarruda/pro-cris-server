package google

import (
	"context"

	"google.golang.org/api/idtoken"
)

type User struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Picture       string `json:"picture"`
	EmailVerified bool   `json:"emailVerified"`
}

type Client struct {
	ClientId  string
	SecretKey string
}

func (c Client) Validate(idToken string) (User, error) {
	payload, err := idtoken.Validate(context.Background(), idToken, c.ClientId)
	if err != nil {
		return User{}, err
	}
	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)
	picture := payload.Claims["picture"].(string)
	emailVerified := payload.Claims["email_verified"].(bool)

	return User{
		Name:          name,
		Email:         email,
		Picture:       picture,
		EmailVerified: emailVerified,
	}, nil
}
