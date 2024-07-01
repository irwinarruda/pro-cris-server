package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
)

func CreateAuthRoutes(app *gin.RouterGroup) {
	authCtrl := proinject.Resolve(&auth.AuthCtrl{})
	app.POST("/v1/auth/login", authCtrl.Login)
}
