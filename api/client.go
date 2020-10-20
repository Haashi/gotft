package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	request, _ := http.NewRequest("GET", fmt.Sprintf(URLFormat, c.region, url), nil)
	request.Header.Add(tokenHeader, c.apiKey)
	res, err := c.c.Do(request)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusTooManyRequests {
		retry := res.Header.Get("Retry-After")
		seconds, _ := strconv.Atoi(retry)
		fmt.Fprintf(os.Stdout, "rate limited, retrying in %ds\n", seconds)
		time.Sleep(time.Duration(seconds) * time.Second)
		return c.Get(url)
	}
	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("ressource not found %s", url)
	}
	if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("unauthorized or forbidden, api key is probably wrong or expired")
	}
	if res.StatusCode == http.StatusBadRequest {
		var data struct {
			Status struct {
				Message    string `json:"message"`
				StatusCode int    `json:"status_code"`
			} `json:"status"`
		}
		json.NewDecoder(res.Body).Decode(&data)
		return nil, fmt.Errorf("%+v", data.Status.Message)
	}
	return res.Body, nil
}
