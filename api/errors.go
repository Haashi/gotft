package api

import "fmt"

type ErrorNotFound struct {
	url string
}

func (err ErrorNotFound) Error() string {
	return fmt.Sprintf("ressource not found : %s", err.url)
}

type ErrorUnauthorized struct {
	url    string
	apiKey string
}

func (err ErrorUnauthorized) Error() string {
	return fmt.Sprintf("unauthorized, api key is wrong or missing or ressouce is not found : %s / %s", err.apiKey, err.url)
}

type ErrorNetwork struct {
	url string
}

func (err ErrorNetwork) Error() string {
	return fmt.Sprintf("network error getting %s", err.url)
}

type ErrorForbidden struct {
	url    string
	apiKey string
}

func (err ErrorForbidden) Error() string {
	return fmt.Sprintf("forbidden, api key is banned or path is wrong : %s / %s", err.apiKey, err.url)
}

type ErrorBadRequest struct {
	url  string
	info string
}

func (err ErrorBadRequest) Error() string {
	return fmt.Sprintf("bad request on %s, info from server : %s", err.url, err.info)
}

type ErrorDecode struct {
	target string
	info   string
}

func (err ErrorDecode) Error() string {
	return fmt.Sprintf("error decoding %s : %s", err.target, err.info)
}
