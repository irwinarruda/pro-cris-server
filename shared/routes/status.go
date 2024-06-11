package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/modules/status"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

func CreateStatusRoutes(app *gin.RouterGroup) {
	statusCtrl := configs.ResolveCtrl(&status.StatusCtrl{})
	app.GET("/status", statusCtrl.GetStatus)
}
