package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/irwinarruda/pro-cris-server/libs/openai"
	"github.com/irwinarruda/pro-cris-server/libs/whatsapp"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/templates"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"google.golang.org/api/idtoken"
)

func chatbotServer() {
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

func templateServer() {
	app := gin.New()
	count := 0
	app.GET("/", func(c *gin.Context) {
		component := templates.App("Some Name", count)
		component.Render(context.Background(), c.Writer)
	})
	app.POST("/increment", func(c *gin.Context) {
		count++
		c.String(http.StatusOK, fmt.Sprintf("Count: %v", count))
	})
	app.POST("/auth", func(c *gin.Context) {
		c.String(http.StatusOK, "Authenticated!")
	})
	key := "randomString"
	maxAge := 86400 * 30
	isProd := false
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd
	gothic.Store = store
	env := configs.GetEnv()
	goth.UseProviders(
		google.New(env.GoogleClientId, env.GoogleSecretKey, "http://localhost:8080/auth/google/callback"),
	)

	app.GET("/auth/:provider/callback", func(c *gin.Context) {
		provider := c.Param("provider")
		request := c.Request.WithContext(context.WithValue(context.Background(), "provider", provider))
		user, err := gothic.CompleteUserAuth(c.Writer, request)
		if err != nil {
			fmt.Fprintln(c.Writer, err)
			return
		}
		fmt.Println(user)
		component := templates.App("User Authenticated", count)
		component.Render(context.Background(), c.Writer)
	})

	app.GET("/logout/{provider}", func(c *gin.Context) {
		// gothic.Logout(res, req)
		// res.Header().Set("Location", "/")
		// res.WriteHeader(http.StatusTemporaryRedirect)
	})

	app.GET("/auth/:provider", func(c *gin.Context) {
		// try to get the user without re-authenticating
		provider := c.Param("provider")
		request := c.Request.WithContext(context.WithValue(context.Background(), "provider", provider))
		if _, err := gothic.CompleteUserAuth(c.Writer, request); err == nil {
			component := templates.App("Authenticated", count)
			component.Render(context.Background(), c.Writer)
		} else {
			gothic.BeginAuthHandler(c.Writer, c.Request)
		}
	})

	app.Run()
}

func androidServer() {
	app := gin.New()
	app.POST("/validateToken", func(c *gin.Context) {
		body := struct {
			Token string `json:"token"`
		}{}
		c.Bind(&body)
		env := configs.GetEnv()
		payload, err := idtoken.Validate(context.Background(), body.Token, env.GoogleClientId)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusForbidden, "")
			return
		}
		fmt.Println("%v", payload)
		c.String(http.StatusOK, "")
	})

	app.Run()
}

func main() {
	// runServer()
	// templateServer()
	androidServer()
}
