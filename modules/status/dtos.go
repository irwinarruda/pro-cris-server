package status

import "time"

type GetStatusDatabaseDTO struct {
	Version         string `json:"version" validate:"required"`
	MaxConnections  int    `json:"max_connections" validate:"number"`
	OpenConnections int    `json:"open_connections" validate:"number"`
}

type GetStatusDependenciesDTO struct {
	Database GetStatusDatabaseDTO `json:"database" validate:"required"`
}

type GetStatusDTO struct {
	UpdatedAt    time.Time                `json:"updated_at" validate:"required"`
	Dependencies GetStatusDependenciesDTO `json:"dependencies" validate:"required"`
}
