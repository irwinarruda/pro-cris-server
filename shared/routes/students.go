package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/irwinarruda/pro-cris-server/modules/students"
)

func CreateStudentRoutes(app *gin.RouterGroup) {
	app.POST("/students", students.CreateStudent)
	app.GET("/students", students.CreateStudent)
}
