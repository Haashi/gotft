package gotft

import (
	"sync"
	"testing"
	"time"

	"github.com/haashi/gotft/internal"
)

func TestNoInternetInternal(t *testing.T) {
	c := newClient(apiKey, "badregion", &Options{c: &internal.TimeoutClient{}, log: testOpt.log})
	_, err := c.get("/")
	if _, ok := err.(ErrorNetwork); !ok {
		t.Errorf("error missing or error is not a network error")
	}
}

func TestNoInternet(t *testing.T) {
	c := newClient(apiKey, "badregion", testOpt)
	_, err := c.get("/league/v1/master")
	if _, ok := err.(ErrorNetwork); !ok {
		t.Errorf("error missing or error is not a network error")
	}
}

func TestUnauthorizedInternal(t *testing.T) {
	c := newClient(apiKey, "badregion", &Options{c: &internal.UnauthorizedClient{}, log: testOpt.log})
	_, err := c.get("/")
	if _, ok := err.(ErrorUnauthorized); !ok {
		t.Errorf("error missing or error is not a unauthorized error")
	}
}

func TestForbiddenInternal(t *testing.T) {
	c := newClient(apiKey, "badregion", &Options{c: &internal.ForbiddenClient{}, log: testOpt.log})
	_, err := c.get("/")
	if _, ok := err.(ErrorForbidden); !ok {
		t.Errorf("error missing or error is not a forbidden error")
	}
}

func TestForbidden(t *testing.T) {
	c := newClient("thisisawrongapikey", EUROPE, testOpt)
	_, err := c.get("/league/v1/master")
	if _, ok := err.(ErrorForbidden); !ok {
		t.Errorf("error missing or error is not a forbidden error")
	}
}

func TestNotFoundInternal(t *testing.T) {
	c := newClient(apiKey, "badregion", &Options{c: &internal.NotFoundClient{}, log: testOpt.log})
	_, err := c.get("/")
	if _, ok := err.(ErrorNotFound); !ok {
		t.Errorf("error missing or error is not a notfound error")
	}
}

func TestNotFound(t *testing.T) {
	c := newClient(apiKey, EUROPE, testOpt)
	_, err := c.get("/match/v1/matches/EUW1_4000230362")
	if _, ok := err.(ErrorNotFound); !ok {
		t.Errorf("error missing or error is not a notfound error")
	}
}

func TestTooManyRequestInternal(t *testing.T) {
	start := time.Now()
	c := newClient(apiKey, "badregion", &Options{c: &internal.TooManyRequestsClient{}, log: testOpt.log})
	_, err := c.get("/")
	if err != nil || time.Since(start) <= time.Second {
		t.Errorf("error too many request, did not wait a minute before second get or failed")
	}
}

func TestTooManyRequest(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	var wg sync.WaitGroup
	for i := 0; i < 22; i++ {
		wg.Add(1)
		go func() {
			_, err := c.get("/league/v1/master")
			if err != nil {
				t.Errorf("failed to get league/v1/master")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
