package configs

import (
	"github.com/irwinarruda/pro-cris-server/libs/proenv"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"github.com/joho/godotenv"
)

type Env struct {
	BaseUrl           string `env:"PRO_CRIS_BASE_URL"`
	Port              string `env:"PRO_CRIS_SERVER_PORT"`
	WhatsAppAuthToken string `env:"API_WHATSAPP_AUTH_TOKEN"`
	WhatsAppUrl       string `env:"API_WHATSAPP_URL"`
	WhatsAppPhoneId   string `env:"API_WHATSAPP_PHONE_ID"`
	WhatsAppChallenge string `env:"API_WHATSAPP_VERIFY_TOKEN"`
	OpenAiUrl         string `env:"API_OPENAI_URL"`
	OpenAiAuthToken   string `env:"API_OPENAI_AUTH_TOKEN"`
	GoogleClientId    string `env:"API_GOOGLE_CLIENT_ID"`
	GoogleSecretKey   string `env:"API_GOOGLE_SECRET_KEY"`
}

var env *Env

func GetEnv(filenames ...string) Env {
	if env == nil {
		err := godotenv.Load(filenames...)
		utils.AssertErr(err)
		env = &Env{}
	}
	err := proenv.LoadEnv(env)
	utils.AssertErr(err)
	return *env
}
