package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/status"
)

func CreateStatusRoutes(app *gin.RouterGroup) {
	statusCtrl := proinject.Resolve(&status.StatusCtrl{})
	app.GET("/v1/status", statusCtrl.GetStatus)
}
