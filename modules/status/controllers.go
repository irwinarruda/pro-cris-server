package status

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

type StatusCtrl struct {
	Env      configs.Env      `ctrl:"env"`
	Validate configs.Validate `ctrl:"validate"`
}

func (s StatusCtrl) GetStatus(c *gin.Context) {
	status := GetStatusDTO{
		UpdatedAt: time.Now(),
		Dependencies: GetStatusDependenciesDTO{
			Database: GetStatusDatabaseDTO{
				Version:         "16",
				MaxConnections:  0,
				OpenConnections: 1,
			},
		},
	}
	s.Validate.Struct(status)
	c.JSON(http.StatusOK, status)
}
