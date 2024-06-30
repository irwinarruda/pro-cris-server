package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

func CreateAuthRoutes(app *gin.RouterGroup) {
	authCtrl := configs.ResolveInject(&auth.AuthCtrl{})
	app.POST("/v1/auth/login", authCtrl.Login)
}
