package prohttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Config[T any] struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    *T
}

func NewRequest[T any](reqConfig Config[T]) (*http.Request, error) {
	jsonBody, err := json.Marshal(reqConfig.Body)
	if err != nil {
		return nil, errors.New("[prohttp]: Could not serialize JSON body")
	}
	bodyReader := bytes.NewReader([]byte(jsonBody))
	req, err := http.NewRequest(reqConfig.Method, reqConfig.Url, bodyReader)
	if err != nil {
		return nil, errors.New("[prohttp]: Request could not be created")
	}
	if len(reqConfig.Headers) != 0 {
		for key, value := range reqConfig.Headers {
			req.Header.Set(key, value)
		}
	}
	return req, nil
}
