package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

func main() {
	env := configs.GetEnv(".env")
	app := gin.New()
	app.GET("/", func(c *gin.Context) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(c.Writer, env)
	})
	app.GET("/redirect", func(c *gin.Context) {
		c.File("./templates/redirect.html")
	})
	app.Run(":3000")
}
