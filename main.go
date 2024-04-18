package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/wa"
	"github.com/irwinarruda/pro-cris-server/modules/whatsapp"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

func runServer() {
	env := configs.GetEnv()
	app := gin.New()
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
		resBody := wa.ResData{}
		err := c.Bind(&resBody)
		if err != nil {
			fmt.Printf("Usuário: %v\n", err)
		}
		value := wa.GetResMessage(&resBody)
		if value != nil {
			fmt.Println((*value).Text.Body)
		}
		c.JSON(http.StatusNoContent, "")
	})
	app.Run()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	go runServer()
	for scanner.Scan() {
		text := scanner.Text()
		if strings.ToLower(text) == "quit" {
			os.Exit(0)
		}
		body, err := wa.NewReqTextMessage("5562982584840", &wa.ReqText{
			Body:       text,
			PreviewUrl: false,
		})

		utils.AssertErr(err)
		err = whatsapp.SendMessage(body)
		utils.AssertErr(err)
	}
}
