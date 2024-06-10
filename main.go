package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/google"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
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
		user, err := goo.Validate(body.Token)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusForbidden, err.Error())
			return
		}
		c.JSON(http.StatusOK, user)
	})
	app.Run()
}

func main() {
	app := gin.New()
	v1 := app.Group("/api/v1")
	routes.CreateRoutes(v1)
	app.Run()
}
