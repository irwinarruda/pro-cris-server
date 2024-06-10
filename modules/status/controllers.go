package status

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
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
	c.JSON(http.StatusOK, status)
}
