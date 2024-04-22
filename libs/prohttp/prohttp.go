package prohttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type RequestConfig[T any] struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    *T
}

type ResponseConfig[S any] struct {
	Url        string
	Method     string
	StatusCode int
	Body       []byte
}

func (r ResponseConfig[S]) IsOk() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

func (r ResponseConfig[S]) ParseBody(body *S) bool {
	err := json.Unmarshal(r.Body, body)
	return err == nil
}

func (r ResponseConfig[S]) RawBody() string {
	return string(r.Body)
}

func DoRequest[S any, T any](reqConfig RequestConfig[T]) (ResponseConfig[S], error) {
	jsonBody, err := json.Marshal(reqConfig.Body)
	if err != nil {
		return ResponseConfig[S]{}, errors.New("[prohttp]: Could not serialize JSON body")
	}
	bodyReader := bytes.NewReader([]byte(jsonBody))
	req, err := http.NewRequest(reqConfig.Method, reqConfig.Url, bodyReader)
	if err != nil {
		return ResponseConfig[S]{}, errors.New("[prohttp]: Request could not be created")
	}
	req.Header.Set("Content-Type", "application/json")
	if len(reqConfig.Headers) != 0 {
		for key, value := range reqConfig.Headers {
			req.Header.Set(key, value)
		}
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return ResponseConfig[S]{}, errors.New("[prohttp]: Could not execute the request")
	}
	defer res.Body.Close()
	byteBody, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseConfig[S]{}, errors.New("[prohttp]: Could not read the body")
	}
	return ResponseConfig[S]{
		Url:        (*res.Request.URL).String(),
		Method:     res.Request.Method,
		StatusCode: res.StatusCode,
		Body:       byteBody,
	}, nil
}
