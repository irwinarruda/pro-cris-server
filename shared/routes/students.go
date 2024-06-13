package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

func CreateStudentRoutes(app *gin.RouterGroup) {
	studentsCtrl := configs.ResolveInject(&students.StudentCtrl{})
	app.POST("/students", studentsCtrl.CreateStudent)
	app.GET("/students", studentsCtrl.CreateStudent)
}
