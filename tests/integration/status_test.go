package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/modules/status"
	"github.com/stretchr/testify/assert"
)

func TestStatusService(t *testing.T) {
	Init()

	t.Run("Happy Path", func(t *testing.T) {
		var assert = assert.New(t)
		var statusService = status.NewStatusService()
		var status = statusService.GetStatus()

		assert.NotEqual(0, status.UpdatedAt, "Should return a valid updated at.")
		assert.NotEqual(0, status.Dependencies.Database, "Should return a valid database status.")
		assert.Equal("16.0", status.Dependencies.Database.Version, "Should return a valid database version.")
		assert.LessOrEqual(0, status.Dependencies.Database.MaxConnections, "Shoud return 0 as max connections.")
		assert.Equal(1, status.Dependencies.Database.OpenConnections, "Should return a 1 as database open connections.")
		assert.LessOrEqual(status.Dependencies.Database.OpenConnections, status.Dependencies.Database.MaxConnections, "Database open connections should be less than or equal to max connections")
	})
}
