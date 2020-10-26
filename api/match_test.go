package api

import (
	"testing"
)

func TestNewMatchClient(t *testing.T) {
	c := newClient(apiKey, EUROPE, testOpt)
	_ = newMatchClient(c, testOpt)
}

func Test_matchClient_GetByPuuid(t *testing.T) {
	c := newClient(apiKey, EUROPE, testOpt)
	mc := newMatchClient(c, testOpt)
	_, err := mc.GetMatchesByPuuid("5_QD37vWUa7Eq8jP6Cy-R18z60E9nRJlpkDZjqBGvtngedjANG6221udHyYnN2wCJCZV7CnlAqcnHQ", 10)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_matchClient_GetByPuuid_Fail(t *testing.T) {
	c := newClient(apiKey, EUROPE, testOpt)
	mc := newMatchClient(c, testOpt)
	_, err := mc.GetMatchesByPuuid("6_QD37vWUa7Eq8jP6Cy-R18z60E9nRJlpkDZjqBGvtngedjANG6221udHyYnN2wCJCZV7CnlAqcnHQ", 10)
	if _, ok := err.(ErrorBadRequest); !ok {
		t.Errorf("error missing or error is not a badrequest error")
	}
}

func Test_matchClient_GetMatch(t *testing.T) {
	c := newClient(apiKey, EUROPE, testOpt)
	mc := newMatchClient(c, testOpt)
	_, err := mc.GetMatch("EUW1_4770230362")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_matchClient_GetMatch_Fail(t *testing.T) {
	c := newClient(apiKey, EUROPE, testOpt)
	mc := newMatchClient(c, testOpt)
	_, err := mc.GetMatch("EUW1_4750230362")
	if _, ok := err.(ErrorNotFound); !ok {
		t.Errorf("error missing or error is not a notfound error")
	}
}
