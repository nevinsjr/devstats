package client

import (
	"devstats/internal/utils"
	"io"
	"net/http"
)

type HttpRequestWrapper struct {
	*http.Request
}

func NewRequest(method string, url string, body io.Reader) (*HttpRequestWrapper, error) {
	request, err := http.NewRequest(method, url, body)
	utils.CheckFatalError(err)

	return &HttpRequestWrapper{request}, nil
}

func (request *HttpRequestWrapper) WithOptions(options ...Option) (*HttpRequestWrapper, error) {
	var err error
	updatedRequest := request

	for _, opt := range options {
		updatedRequest, err = opt(updatedRequest)
		if err != nil {
			return updatedRequest, err
		}
	}

	return updatedRequest, nil
}

type Option func(request *HttpRequestWrapper) (*HttpRequestWrapper, error)

func BasicAuth(username string, password string) Option {
	return func(request *HttpRequestWrapper) (*HttpRequestWrapper, error) {
		request.SetBasicAuth(username, password)
		return request, nil
	}
}

func Header(query string, args string) Option {
	return func(request *HttpRequestWrapper) (*HttpRequestWrapper, error) {
		request.Header.Add(query, args)
		return request, nil
	}
}

func QueryParam(key string, value string) Option {
	return func(request *HttpRequestWrapper) (*HttpRequestWrapper, error) {
		query := request.URL.Query()
		query.Add(key, value)
		request.URL.RawQuery = query.Encode()
		return request, nil
	}
}
