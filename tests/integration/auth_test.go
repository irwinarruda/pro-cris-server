package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/stretchr/testify/assert"
)

func TestAuthService(t *testing.T) {
	setupTestsAuth()
	var assert = assert.New(t)
	var _ = auth.NewAuthService()
	assert.Equal(0, 0)
}

func setupTestsAuth() {
	configs.RegisterInject("env", configs.GetEnv("../../.env"))
	configs.RegisterInject("db", configs.GetDb())
}
