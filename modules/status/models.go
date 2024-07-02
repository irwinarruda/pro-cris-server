package status

import "time"

type StatusDatabase struct {
	Version         string `json:"version" validate:"required"`
	MaxConnections  int    `json:"maxConnections"`
	OpenConnections int    `json:"openConnections"`
}

type StatusDependencies struct {
	Database StatusDatabase `json:"database" validate:"required"`
}

type Status struct {
	UpdatedAt    time.Time          `json:"updatedAt" validate:"required"`
	Dependencies StatusDependencies `json:"dependencies" validate:"required"`
}
