package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/wa"
	"github.com/irwinarruda/pro-cris-server/modules/chatbot"
	"github.com/irwinarruda/pro-cris-server/modules/whatsapp"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

func runServer() {
	env := configs.GetEnv()
	app := gin.New()
	doPrompts := false
	app.GET("/message", func(c *gin.Context) {
		token := c.Query("hub.verify_token")
		challenge := c.Query("hub.challenge")
		if token == env.WhatsAppChallenge {
			c.String(http.StatusOK, challenge)
			return
		}
		c.String(http.StatusBadRequest, "")
	})
	app.POST("/message", func(c *gin.Context) {
		resBody := wa.ResInfo{}
		if err := c.Bind(&resBody); err != nil {
			fmt.Printf("%v\n", err)
			c.JSON(http.StatusNoContent, "")
			return
		}
		if value, ok := wa.GetResMessage(&resBody); ok {
			switch strings.ToLower(value.Text.Body) {
			case strings.ToLower("Iniciar GPT"):
				body := wa.NewReqTextMessage("5562982584840", "Parabéns, você iniciou o modo GPT, faça uma pergunta")
				whatsapp.SendMessage(&body)
				doPrompts = true
			case strings.ToLower("Sair"):
				doPrompts = false
			default:
				if doPrompts {
					body := wa.NewReqTextMessage("5562982584840", "Pensando...")
					err := whatsapp.SendMessage(&body)
					if err != nil {
						c.JSON(http.StatusNoContent, "")
						return
					}
					resChat := chatbot.SendPrompt(value.Text.Body)
					body = wa.NewReqTextMessage("5562982584840", resChat.Choices[0].Message.Content)
					err = whatsapp.SendMessage(&body)
					if err != nil {
						c.JSON(http.StatusNoContent, "")
						return
					}
				}
			}
		}
		c.JSON(http.StatusNoContent, "")
	})
	app.Run()
}

func main() {
	runServer()
}
