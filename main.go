package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/google"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/containers"
	"github.com/irwinarruda/pro-cris-server/shared/routes"
)

func androidServer() {
	env := configs.GetEnv()
	goo := google.Client{
		ClientId:  env.GoogleClientId,
		SecretKey: env.GoogleSecretKey,
	}
	app := gin.New()
	app.POST("/validateToken", func(c *gin.Context) {
		body := struct {
			Token string `json:"token"`
		}{}
		c.Bind(&body)
		account, err := goo.Validate(body.Token)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusForbidden, err.Error())
			return
		}
		c.JSON(http.StatusOK, account)
	})
	app.Run()
}

func main() {
	containers.InitInjections()
	app := gin.New()
	v1 := app.Group("/api")
	routes.CreateApiRoutes(v1)
	app.Run()
}
