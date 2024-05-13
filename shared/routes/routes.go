package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/irwinarruda/pro-cris-server/modules/students"
)

func CreateRoutes(app *gin.Engine) {
	app.POST("/students", students.CreateStudent)
}
