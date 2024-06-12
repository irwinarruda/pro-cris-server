package status

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

type StatusCtrl struct {
	Validate configs.Validate `ctrl:"validate"`
	Db       configs.Db       `ctrl:"db"`
}

func (s StatusCtrl) GetStatus(c *gin.Context) {
	status := GetStatusDTO{
		UpdatedAt: time.Now(),
		Dependencies: GetStatusDependenciesDTO{
			Database: GetStatusDatabaseDTO{},
		},
	}
	s.Validate.Struct(status)
	c.JSON(http.StatusOK, status)
}
