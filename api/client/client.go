package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const URLFormat = "https://%s.api.riotgames.com/tft%s"
const tokenHeader = "X-Riot-Token"

const (
	BR  = "br"
	EUW = "euw"
	EUN = "eun"
	JP  = "jp"
	KR  = "kr"
	LAN = "lan"
	LAS = "las"
	NA  = "na"
	OCE = "oce"
	TR  = "tr"
	RU  = "ru"
)

var globalRouting = map[string]string{
	BR:  "americas",
	EUW: "europe",
	EUN: "europe",
	JP:  "asia",
	KR:  "asia",
	LAN: "americas",
	LAS: "americas",
	NA:  "americas",
	OCE: "asia",
	TR:  "europe",
	RU:  "europe",
}

var routing = map[string]string{
	BR:  "br1",
	EUW: "euw1",
	EUN: "eun1",
	JP:  "jp1",
	KR:  "kr",
	LAN: "la1",
	LAS: "la2",
	NA:  "na1",
	OCE: "oc1",
	TR:  "tr1",
	RU:  "ru",
}

type Client struct {
	c        *http.Client
	limiters []*rate.Limiter
	region   string
	apiKey   string
}

func NewClient(apikey string, region string, prodKey bool) *Client {
	c := Client{c: &http.Client{}, region: region, apiKey: apikey}
	//TODO : prodkey has better rate limiting
	rt1 := rate.Every(time.Second / 20)      // rate 20 calls per second
	rt2 := rate.Every(2 * time.Minute / 100) // rate 100 calls per 2 minutes
	c.limiters = make([]*rate.Limiter, 0)
	c.limiters = append(c.limiters, rate.NewLimiter(rt1, 20))
	c.limiters = append(c.limiters, rate.NewLimiter(rt2, 100))
	return &c
}

func (c *Client) Get(url string, global bool) (io.ReadCloser, error) {
	for _, limiter := range c.limiters {
		limiter.Wait(context.Background())
	}
	var region string
	if global {
		region = globalRouting[c.region]
	} else {
		region = routing[c.region]
	}
	request, err := http.NewRequest("GET", fmt.Sprintf(URLFormat, region, url), nil)
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
