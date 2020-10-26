package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/haashi/gotft/api/internal"
)

const URLFormat = "https://%s.api.riotgames.com/tft%s"
const tokenHeader = "X-Riot-Token"

type apiclient struct {
	c      internal.HttpClient
	region region
	apiKey string
	log    internal.Logger
}

func newClient(apiKey string, region region, opt *Options) *apiclient {
	opt.log.Debug("initializing new http client")
	apiC := apiclient{c: opt.c, region: region, apiKey: apiKey, log: opt.log}
	return &apiC
}

func (c *apiclient) get(url string) (io.ReadCloser, error) {
	request, _ := http.NewRequest("GET", fmt.Sprintf(URLFormat, c.region, url), nil)
	c.log.Debug("getting " + request.URL.String())
	request.Header.Add(tokenHeader, c.apiKey)
	res, err := c.c.Do(request)
	if err != nil {
		c.log.Errorf("error getting %s : %s", request.URL.String(), err.Error())
		return nil, ErrorNetwork{url}
	}
	if res.StatusCode == http.StatusTooManyRequests {
		retry := res.Header.Get("Retry-After")
		seconds, _ := strconv.Atoi(retry)
		c.log.Infof("rate limited, retrying in %ds", seconds)
		time.Sleep(time.Duration(seconds) * time.Second)
		return c.get(url)
	}

	if res.StatusCode == http.StatusNotFound {
		c.log.Errorf("ressource not found %s", url)
		return nil, ErrorNotFound{url}
	}

	if res.StatusCode == http.StatusUnauthorized {
		c.log.Errorf("unauthorized, api key is wrong or missing")
		return nil, ErrorUnauthorized{c.apiKey, url}
	}

	if res.StatusCode == http.StatusForbidden {
		c.log.Errorf("forbidden, api key is banned or path is wrong : %s", c.apiKey)
		return nil, ErrorForbidden{c.apiKey, request.URL.String()}
	}

	if res.StatusCode == http.StatusBadRequest {
		var data struct {
			Status struct {
				Message    string `json:"message"`
				StatusCode int    `json:"status_code"`
			} `json:"status"`
		}
		err := json.NewDecoder(res.Body).Decode(&data)
		if err != nil {

			c.log.Errorf("error decoding status bad request : %s", c.apiKey)
			return nil, ErrorDecode{"status bad request", err.Error()}
		}

		c.log.Errorf("bad request on %s, info from server : %s", url, data.Status.Message)
		return nil, ErrorBadRequest{url, data.Status.Message}
	}
	return res.Body, nil
}
