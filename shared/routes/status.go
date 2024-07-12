package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/modules/status/drivers"
)

func CreateStatusRoutes(app *gin.RouterGroup) {
	var statusCtrl = statusdrivers.NewStatusCtrl()
	app.GET("/v1/status", statusCtrl.GetStatus)
}
