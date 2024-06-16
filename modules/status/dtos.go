package status

import "time"

type GetStatusDatabaseDTO struct {
	Version         string `json:"version" validate:"required"`
	MaxConnections  int    `json:"maxConnections"`
	OpenConnections int    `json:"openConnections"`
}

type GetStatusDependenciesDTO struct {
	Database GetStatusDatabaseDTO `json:"database" validate:"required"`
}

type GetStatusDTO struct {
	UpdatedAt    time.Time                `json:"updatedAt" validate:"required"`
	Dependencies GetStatusDependenciesDTO `json:"dependencies" validate:"required"`
}
