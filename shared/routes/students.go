package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

func CreateStudentRoutes(app *gin.RouterGroup) {
	studentsCtrl := configs.ResolveInject(&students.StudentCtrl{})
	app.GET("/v1/students", studentsCtrl.GetStudents)
	app.GET("/v1/students/:id", studentsCtrl.GetStudent)
	app.POST("/v1/students", studentsCtrl.CreateStudent)
	app.PUT("/v1/students/:id", studentsCtrl.UpdateSudent)
	app.DELETE("/v1/students/:id", studentsCtrl.DeleteStudent)
}
