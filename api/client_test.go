package api

import (
	"sync"
	"testing"
)

func TestNoInternet(t *testing.T) {
	c := NewClient(apiKey, "badregion")
	_, err := c.Get("/league/v1/master")
	if err == nil {
		t.FailNow()
	}
}

func TestWrongApiKey(t *testing.T) {
	c := NewClient("thisisawrongapikey", EUROPE)
	_, err := c.Get("/league/v1/master")
	if err == nil {
		t.FailNow()
	}
}

func TestNotFound(t *testing.T) {
	c := NewClient(apiKey, EUROPE)
	_, err := c.Get("/match/v1/matches/EUW1_4000230362")
	if err == nil {
		t.FailNow()
	}
}

func TestTooManyRequest(t *testing.T) {
	c := NewClient(apiKey, EUW)
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
