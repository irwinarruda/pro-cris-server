package status

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

type StatusRepository struct {
	Db  configs.Db  `inject:"db"`
	Env configs.Env `inject:"env"`
}

func NewStatusRepository() *StatusRepository {
	return proinject.Resolve(&StatusRepository{})
}

func (s *StatusRepository) GetStatusDatabase() StatusDatabase {
	databaseResults := struct {
		ServerVersion  string
		MaxConnections int
		Count          int
	}{}
	s.Db.Raw("SHOW SERVER_VERSION;").Scan(&databaseResults)
	s.Db.Raw("SHOW MAX_CONNECTIONS;").Scan(&databaseResults)
	s.Db.Raw("SELECT COUNT(*) FROM PG_STAT_ACTIVITY WHERE datname = ?;", s.Env.DatabaseName).Scan(&databaseResults)

	return StatusDatabase{
		Version:         databaseResults.ServerVersion,
		MaxConnections:  databaseResults.MaxConnections,
		OpenConnections: databaseResults.Count,
	}
}
