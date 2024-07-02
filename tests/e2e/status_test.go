package e2e

import (
	"testing"
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/prohttp"
	"github.com/irwinarruda/pro-cris-server/modules/status"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/stretchr/testify/assert"
)

func TestGetReturnOK(t *testing.T) {
	time.Sleep(1 * time.Second)
	var assert = assert.New(t)
	var env = configs.GetEnv("../../.env")
	res, err := prohttp.DoRequest[status.Status](prohttp.RequestConfig[any]{
		Url:    env.BaseUrl + "/api/v1/status",
		Method: "GET",
	})
	assert.NoError(err, "GET should not throw an error")
	assert.True(res.IsOk(), "GET should return good status codes")
	body := status.Status{}
	ok := res.ParseBody(&body)
	assert.True(ok, "GET should return a json body")
	database := body.Dependencies.Database
	assert.Equal("16.0", database.Version, "Database version should be 16.0")
	assert.LessOrEqual(0, database.MaxConnections, "Database max connections should greater than or equal to 0")
	assert.Equal(1, database.OpenConnections, "Database open connections should be 1")
	assert.LessOrEqual(database.OpenConnections, database.MaxConnections, "Database open connections should be less than or equal to max connections")
}
