package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/modules/auth/drivers"
)

func CreateAuthRoutes(app *gin.RouterGroup) {
	var authCtrl = authdrivers.NewAuthCtrl()
	app.POST("/v1/auth/login", authCtrl.Login)
}
