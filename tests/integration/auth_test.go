package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/stretchr/testify/assert"
)

func TestAuthService(t *testing.T) {
	configs.GetEnv("../../.env")
	var assert = assert.New(t)
	var authService = auth.NewAuthService()
	id := authService.Login(struct {
		Provider string
		Token    string
	}{
		Provider: "google",
		Token:    "123456",
	})
	_, err := authService.GetUserByID(id)
	assert.NoError(err, "User should be found after first login")
	assert.Equal(id, 0)
}
