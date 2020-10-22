package api

import (
	"testing"
)

func TestNewMatchClient(t *testing.T) {
	c := NewClient(apiKey, EUROPE, log)
	_ = NewMatchClient(c, log)
}

func Test_matchClient_GetByPuuid(t *testing.T) {
	c := NewClient(apiKey, EUROPE, log)
	mc := NewMatchClient(c, log)
	_, err := mc.GetMatchesByPuuid("5_QD37vWUa7Eq8jP6Cy-R18z60E9nRJlpkDZjqBGvtngedjANG6221udHyYnN2wCJCZV7CnlAqcnHQ", 10)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_matchClient_GetByPuuid_Fail(t *testing.T) {
	c := NewClient(apiKey, EUROPE, log)
	mc := NewMatchClient(c, log)
	_, err := mc.GetMatchesByPuuid("6_QD37vWUa7Eq8jP6Cy-R18z60E9nRJlpkDZjqBGvtngedjANG6221udHyYnN2wCJCZV7CnlAqcnHQ", 10)
	if err == nil || err.Code != ErrorUnauthorized {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorUnauthorized)
	}
}

func Test_matchClient_GetMatch(t *testing.T) {
	c := NewClient(apiKey, EUROPE, log)
	mc := NewMatchClient(c, log)
	_, err := mc.GetMatch("EUW1_4770230362")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_matchClient_GetMatch_Fail(t *testing.T) {
	c := NewClient(apiKey, EUROPE, log)
	mc := NewMatchClient(c, log)
	_, err := mc.GetMatch("EUW1_4750230362")
	if err == nil || err.Code != ErrorNotFound {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorNotFound)
	}
}
