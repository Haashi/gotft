package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const URLFormat = "https://%s.api.riotgames.com/tft%s"
const tokenHeader = "X-Riot-Token"

type client struct {
	c      *http.Client
	region string
	apiKey string
}

func NewClient(apikey string, region string) *client {
	c := client{c: &http.Client{}, region: region, apiKey: apikey}
	return &c
}

func (c *client) Get(url string) (io.ReadCloser, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf(URLFormat, c.region, url), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(tokenHeader, c.apiKey)
	res, err := c.c.Do(request)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusTooManyRequests {
		retry := res.Header.Get("Retry-After")
		seconds, err := strconv.Atoi(retry)
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(seconds) * time.Second)
		return c.Get(url)
	}
	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("ressource not found")
	}
	return res.Body, nil
}
