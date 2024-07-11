package appointments

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
)

type AppointmentCtrl struct{}

func NewAppointmentCtrl() *AppointmentCtrl {
	return proinject.Resolve(&AppointmentCtrl{})
}

func (a *AppointmentCtrl) CreateAppointment(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}
