package settlementsdrivers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
)

type SettlementCtrl struct{}

func NewSettlementCtrl() *SettlementCtrl {
	return proinject.Resolve(&SettlementCtrl{})
}

func (a *SettlementCtrl) CreateSettlement(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}
