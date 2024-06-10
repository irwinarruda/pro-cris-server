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
		Url:     env.BaseUrl + "/api/v1/status",
		Method:  "GET",
		Headers: nil,
		Body:    nil,
	})
	if err != nil {
		t.Fatalf("GET should not throw an error")
	}
	if !res.IsOk() {
		t.Fatalf("GET should return good status codes")
	}
	body := status.GetStatusDTO{}
	ok := res.ParseBody(&body)
	if !ok {
		t.Fatalf("GET should return a json body")
	}
	err = validate.Struct(body)
	if err != nil {
		t.Fatalf("GET should return a valid body %v", err)
	}

	database := body.Dependencies.Database
	if database.Version != "16" {
		t.Fatalf("Database version should be 16")
	}
	if database.MaxConnections < 0 {
		t.Fatalf("Database max connections should greater than or equal to 0")
	}
	if database.OpenConnections != 1 {
		t.Fatalf("Database max connections should be 1")
	}
	if database.OpenConnections <= database.MaxConnections {
		t.Fatalf("Database open connections should be less than or equal to max connections")
	}
}
