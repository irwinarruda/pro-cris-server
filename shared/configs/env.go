package configs

import (
	"github.com/irwinarruda/pro-cris-server/libs/proenv"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type Env struct {
	WhatsAppAuthToken string `env:"API_WHATSAPP_AUTH_TOKEN"`
	WhatsAppUrl       string `env:"API_WHATSAPP_URL"`
	WhatsAppPhoneId   string `env:"API_WHATSAPP_PHONE_ID"`
	WhatsAppChallenge string `env:"API_WHATSAPP_VERIFY_TOKEN"`
}

var env *Env

func GetEnv() Env {
	if env == nil {
		env = &Env{}
	}
	err := proenv.LoadEnv(env)
	utils.AssertErr(err)
	return *env
}
