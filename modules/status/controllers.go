package status

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

type StatusCtrl struct {
	Db  configs.Db  `inject:"db"`
	Env configs.Env `inject:"env"`
}

func (s StatusCtrl) GetStatus(c *gin.Context) {
	databaseResults := struct {
		ServerVersion  string
		MaxConnections int
		Count          int
	}{}
	s.Db.Raw("SHOW SERVER_VERSION;").Scan(&databaseResults)
	s.Db.Raw("SHOW MAX_CONNECTIONS;").Scan(&databaseResults)
	s.Db.Raw("SELECT COUNT(*) FROM PG_STAT_ACTIVITY WHERE datname = ?;", s.Env.DatabaseName).Scan(&databaseResults)

	status := GetStatusDTO{
		UpdatedAt: time.Now(),
		Dependencies: GetStatusDependenciesDTO{
			Database: GetStatusDatabaseDTO{
				Version:         databaseResults.ServerVersion,
				MaxConnections:  databaseResults.MaxConnections,
				OpenConnections: databaseResults.Count,
			},
		},
	}
	c.JSON(http.StatusOK, status)
}
