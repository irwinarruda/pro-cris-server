package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/modules/status"
)

func CreateStatusRoutes(app *gin.RouterGroup) {
	var statusCtrl = status.NewStatusCtrl()
	app.GET("/v1/status", statusCtrl.GetStatus)
}
