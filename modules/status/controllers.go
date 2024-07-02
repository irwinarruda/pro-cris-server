package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
)

type StatusCtrl struct{}

func NewStatusCtrl() *StatusCtrl {
	return proinject.Resolve(&StatusCtrl{})
}

func (s StatusCtrl) GetStatus(c *gin.Context) {
	statusService := NewStatusService()
	status := statusService.GetStatus()
	c.JSON(http.StatusOK, status)
}
