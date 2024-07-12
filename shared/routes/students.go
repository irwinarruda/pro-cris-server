package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/modules/auth/drivers"
	"github.com/irwinarruda/pro-cris-server/modules/students/drivers"
)

func CreateStudentRoutes(app *gin.RouterGroup) {
	var studentsCtrl = studentsdrivers.NewStudentCtrl()
	var authCtrl = authdrivers.NewAuthCtrl()
	app.GET("/v1/students", authCtrl.EnsureAuthenticated, studentsCtrl.GetStudents)
	app.GET("/v1/students/:id", authCtrl.EnsureAuthenticated, studentsCtrl.GetStudent)
	app.POST("/v1/students", authCtrl.EnsureAuthenticated, studentsCtrl.CreateStudent)
	app.PUT("/v1/students/:id", authCtrl.EnsureAuthenticated, studentsCtrl.UpdateSudent)
	app.DELETE("/v1/students/:id", authCtrl.EnsureAuthenticated, studentsCtrl.DeleteStudent)
}
