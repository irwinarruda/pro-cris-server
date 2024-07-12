package statusdrivers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/status"
)

type StatusCtrl struct{}

func NewStatusCtrl() *StatusCtrl {
	return proinject.Resolve(&StatusCtrl{})
}

func (s StatusCtrl) GetStatus(c *gin.Context) {
	statusService := status.NewStatusService()
	status := statusService.GetStatus()
	c.JSON(http.StatusOK, status)
}
