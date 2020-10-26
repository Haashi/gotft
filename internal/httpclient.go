package internal

import (
	"fmt"
	"net/http"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func NewDefaultClient() HttpClient {
	return &http.Client{}
}

type TimeoutClient struct {
}

func (c TimeoutClient) Do(r *http.Request) (*http.Response, error) {
	return nil, http.ErrHandlerTimeout
}

type NotFoundClient struct {
}

func (c NotFoundClient) Do(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusNotFound}, nil
}

type UnauthorizedClient struct {
}

func (c UnauthorizedClient) Do(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusUnauthorized}, nil
}

type ForbiddenClient struct {
}

func (c ForbiddenClient) Do(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusForbidden}, nil
}

type TooManyRequestsClient struct {
	count int
}

func (c *TooManyRequestsClient) Do(r *http.Request) (*http.Response, error) {
	fmt.Println(c.count)
	if c.count == 0 {
		c.count += 1
		headers := http.Header{}
		headers.Add("Retry-After", "1")
		return &http.Response{StatusCode: http.StatusTooManyRequests, Header: headers}, nil
	} else {
		return &http.Response{StatusCode: http.StatusOK}, nil
	}
}
