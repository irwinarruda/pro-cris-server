package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

func CreateStudentRoutes(app *gin.RouterGroup) {
	studentsCtrl := configs.ResolveInject(&students.StudentCtrl{})
	app.GET("/students", studentsCtrl.GetStudents)
	app.GET("/students/:id", studentsCtrl.GetStudent)
	app.POST("/students", studentsCtrl.CreateStudent)
	app.PUT("/students/:id", studentsCtrl.UpdateSudent)
}
