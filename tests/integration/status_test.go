package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/prohttp"
	"github.com/irwinarruda/pro-cris-server/modules/status"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

func TestGetReturnOK(t *testing.T) {
	var env = configs.GetEnv("../../.env")
	var validate = configs.GetValidate()
	res, err := prohttp.DoRequest[status.GetStatusDTO](prohttp.RequestConfig[any]{
		Url:    env.BaseUrl + "/api/v1/status",
		Method: "GET",
	})
	if err != nil {
		t.Fatal("GET should not throw an error")
	}
	if !res.IsOk() {
		t.Fatal("GET should return good status codes")
	}
	body := status.GetStatusDTO{}
	ok := res.ParseBody(&body)
	if !ok {
		t.Fatal("GET should return a json body")
	}
	err = validate.Struct(body)
	if err != nil {
		t.Fatalf("GET should return a valid body %v", err)
	}

	database := body.Dependencies.Database
	if database.Version != "16" {
		t.Fatal("Database version should be 16")
	}
	if database.MaxConnections < 0 {
		t.Fatal("Database max connections should greater than or equal to 0")
	}
	if database.OpenConnections != 1 {
		t.Fatal("Database max connections should be 1")
	}
	if database.OpenConnections <= database.MaxConnections {
		t.Fatal("Database open connections should be less than or equal to max connections")
	}
}
