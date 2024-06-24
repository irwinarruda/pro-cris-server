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
	assert := assert.New(t)
	var env = configs.GetEnv("../../.env")
	var validate = configs.GetValidate()
	res, err := prohttp.DoRequest[status.GetStatusDTO](prohttp.RequestConfig[any]{
		Url:    env.BaseUrl + "/api/v1/status",
		Method: "GET",
	})
	assert.NoError(err, "GET should not throw an error")
	assert.True(res.IsOk(), "GET should return good status codes")
	body := status.GetStatusDTO{}
	ok := res.ParseBody(&body)
	assert.True(ok, "GET should return a json body")
	err = validate.Struct(body)
	assert.NoError(err, "GET should return a valid body")
	database := body.Dependencies.Database
	assert.Equal(database.Version, "16.0", "Database version should be 16.0")
	assert.LessOrEqual(0, database.MaxConnections, "Database max connections should greater than or equal to 0")
	assert.Equal(database.OpenConnections, 1, "Database open connections should be 1")
	assert.LessOrEqual(database.OpenConnections, database.MaxConnections, "Database open connections should be less than or equal to max connections")
}
