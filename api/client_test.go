package api

import (
	"sync"
	"testing"
)

func TestNoInternet(t *testing.T) {
	c := NewClient(apiKey, "badregion", log)
	_, err := c.Get("/league/v1/master")
	if err == nil || err.Code != ErrorNetwork {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorNetwork)
	}
}

func TestWrongApiKey(t *testing.T) {
	c := NewClient("thisisawrongapikey", EUROPE, log)
	_, err := c.Get("/league/v1/master")
	if err == nil || err.Code != ErrorUnauthorized {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorUnauthorized)
	}
}

func TestNotFound(t *testing.T) {
	c := NewClient(apiKey, EUROPE, log)
	_, err := c.Get("/match/v1/matches/EUW1_4000230362")
	if err == nil || err.Code != ErrorNotFound {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorNotFound)
	}
}

func TestTooManyRequest(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	var wg sync.WaitGroup
	for i := 0; i < 22; i++ {
		wg.Add(1)
		go func() {
			c.Get("/league/v1/master")
			wg.Done()
		}()
	}
	wg.Wait()
}
