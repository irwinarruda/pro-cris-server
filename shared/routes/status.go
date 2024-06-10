package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/modules/status"
)

func CreateStatusRoutes(app *gin.RouterGroup) {
	app.GET("/status", status.GetStatus)
}
