package providers

import (
	"github.com/irwinarruda/pro-cris-server/libs/google"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

type IGoogleUser = google.User

type IGoogle interface {
	Validate(token string) (IGoogleUser, error)
}

type GoogleClient struct {
	Env    configs.Env `inject:"env"`
	client *google.Client
}

func NewGoogleClient() *GoogleClient {
	googleClient := configs.ResolveInject(&GoogleClient{})
	googleClient.Init()
	return googleClient
}

func (g *GoogleClient) Init() {
	g.client = &google.Client{
		ClientId:  g.Env.GoogleClientId,
		SecretKey: g.Env.GoogleSecretKey,
	}
}

func (g *GoogleClient) Validate(token string) (IGoogleUser, error) {
	user, err := g.client.Validate(token)
	if err != nil {
		return google.User{}, err
	}
	return user, nil
}
