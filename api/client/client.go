package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const URLFormat = "https://%s.api.riotgames.com%s"
const tokenHeader = "X-Riot-Token"

type Client struct {
	c       *http.Client
	limiter *rate.Limiter
	region  string
	apiKey  string
}

func NewClient(apikey string, region string, prodKey bool) *Client {
	c := Client{c: &http.Client{}, region: region, apiKey: apikey}
	//TODO : prodkey has better rate limiting
	rt := rate.Every(2 * time.Minute / 100)
	c.limiter = rate.NewLimiter(rt, 100)
	return &c
}

func (c *Client) Get(url string) (io.ReadCloser, error) {
	c.limiter.Wait(context.Background())
	request, err := http.NewRequest("GET", fmt.Sprintf(URLFormat, c.region, url), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(tokenHeader, c.apiKey)
	res, err := c.c.Do(request)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
