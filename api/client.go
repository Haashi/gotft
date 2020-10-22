package api

import (
	"encoding/json"
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
	region region
	apiKey string
	log    logger
}

func NewClient(apikey string, region region, log logger) *client {
	log.Debug("initializing new http client")
	c := client{c: &http.Client{}, region: region, apiKey: apikey, log: log}
	return &c
}

func (c *client) Get(url string) (io.ReadCloser, *error) {
	request, _ := http.NewRequest("GET", fmt.Sprintf(URLFormat, c.region, url), nil)
	c.log.Debug("getting " + request.URL.String())
	request.Header.Add(tokenHeader, c.apiKey)
	res, err := c.c.Do(request)
	if err != nil {
		c.log.Error("error getting "+request.URL.String(), err.Error())
		return nil, &error{ErrorNetwork, err.Error()}
	}
	if res.StatusCode == http.StatusTooManyRequests {
		retry := res.Header.Get("Retry-After")
		seconds, _ := strconv.Atoi(retry)
		c.log.Infof("rate limited, retrying in %ds\n", seconds)
		time.Sleep(time.Duration(seconds) * time.Second)
		return c.Get(url)
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, &error{ErrorNotFound, fmt.Sprintf("ressource not found %s", url)}
	}

	if res.StatusCode == http.StatusUnauthorized {
		return nil, &error{ErrorUnauthorized, "unauthorized, api key is missing"}
	}

	if res.StatusCode == http.StatusForbidden {
		return nil, &error{ErrorUnauthorized, fmt.Sprintf("unauthorized, api key is wrong or expired : %s", c.apiKey)}
	}

	if res.StatusCode == http.StatusBadRequest {
		var data struct {
			Status struct {
				Message    string `json:"message"`
				StatusCode int    `json:"status_code"`
			} `json:"status"`
		}
		json.NewDecoder(res.Body).Decode(&data)
		return nil, &error{ErrorUnauthorized, fmt.Sprintf("bad request, info from server : %+v", data)}
	}
	return res.Body, nil
}
