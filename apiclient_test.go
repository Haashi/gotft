package gotft

import (
	"testing"
	"time"

	"github.com/haashi/gotft/internal"
)

func TestNoInternet(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.TimeoutClient{}, log: testOpt.log})
	_, err := c.get("/")
	if _, ok := err.(ErrorNetwork); !ok {
		t.Errorf("error missing or error is not a network error")
	}
}

func TestUnauthorized(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.UnauthorizedClient{}, log: testOpt.log})
	_, err := c.get("/")
	if _, ok := err.(ErrorUnauthorized); !ok {
		t.Errorf("error missing or error is not a unauthorized error")
	}
}

func TestForbidden(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.ForbiddenClient{}, log: testOpt.log})
	_, err := c.get("/")
	if _, ok := err.(ErrorForbidden); !ok {
		t.Errorf("error missing or error is not a forbidden error")
	}
}

func TestNotFound(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.NotFoundClient{}, log: testOpt.log})
	_, err := c.get("/")
	if _, ok := err.(ErrorNotFound); !ok {
		t.Errorf("error missing or error is not a notfound error")
	}
}

func TestBadRequest(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadRequestClient{}, log: testOpt.log})
	_, err := c.get("/")
	if _, ok := err.(ErrorBadRequest); !ok {
		t.Errorf("error missing or error is not a badrequest error")
	}
}

func TestTooManyRequest(t *testing.T) {
	start := time.Now()
	c := newClient(apiKey, EUW, &Options{c: &internal.TooManyRequestsClient{}, log: testOpt.log})
	_, err := c.get("/")
	if err != nil || time.Since(start) <= time.Second {
		t.Errorf("error too many request, did not wait a minute before second get or failed")
	}
}

func TestBadRequestBadDecode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadRequestBadDecodeClient{}, log: testOpt.log})
	_, err := c.get("/")
	if _, ok := err.(ErrorDecode); !ok {
		t.Errorf("error missing or error is not a decode error")
	}
}
