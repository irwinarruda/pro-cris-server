package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/openai"
	"github.com/irwinarruda/pro-cris-server/libs/whatsapp"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

func main() {
	env := configs.GetEnv()
	wa := whatsapp.Client{
		Url:       env.WhatsAppUrl,
		PhoneId:   env.WhatsAppPhoneId,
		AuthToken: env.WhatsAppAuthToken,
		Challenge: env.WhatsAppChallenge,
	}
	oa := openai.Client{
		Url:       env.OpenAiUrl,
		AuthToken: env.OpenAiAuthToken,
		Model:     openai.AiModelGpt35Turbo,
		System: []string{
			"You are a personal trainter with focus on food science",
			"Respond with small messages",
		},
	}
	doPrompts := false

	app := gin.New()

	app.GET("/message", func(c *gin.Context) {
		if challenge, ok := wa.ValidateWAWebhookGin(c); ok {
			c.String(http.StatusOK, challenge)
			return
		}
		c.String(http.StatusNotFound, "")
	})

	app.POST("/message", func(c *gin.Context) {
		resMessage, err := wa.GetResMessageGin(c)
		if err != nil {
			c.String(http.StatusNoContent, "")
			return
		}
		switch strings.ToLower(resMessage.Text.Body) {
		case strings.ToLower("Sair"):
			doPrompts = false
		case strings.ToLower("Iniciar GPT"):
			body := whatsapp.NewReqTextMessage("5562982584840", "Parabéns, você iniciou o modo GPT, faça uma pergunta")
			err := wa.SendMessage(&body)
			if err != nil {
				fmt.Println(err.Error())
				c.String(http.StatusNoContent, "")
				return
			}
			doPrompts = true
		default:
			if !doPrompts {
				body := whatsapp.NewReqTextMessage("5562982584840", "Selecione uma opção:\n1. Iniciar GPT")
				wa.SendMessage(&body)
				c.String(http.StatusNoContent, "")
				return
			}
			go func() {
				body := whatsapp.NewReqTextMessage("5562982584840", "Pensando...")
				err := wa.SendMessage(&body)
				if err != nil {
					fmt.Println(err.Error())
					c.String(http.StatusNoContent, "")
					return
				}
			}()
			resChat, err := oa.SendPrompt(resMessage.Text.Body)
			if err != nil {
				fmt.Println(err.Error())
				c.String(http.StatusNoContent, "")
				return
			}
			body := whatsapp.NewReqTextMessage("5562982584840", resChat.Choices[0].Message.Content)
			err = wa.SendMessage(&body)
			if err != nil {
				fmt.Println(err.Error())
				c.String(http.StatusNoContent, "")
				return
			}
		}
		c.String(http.StatusNoContent, "")
		return
	})

	app.Run()
}
