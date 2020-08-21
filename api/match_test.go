package api

import (
	"testing"
)

func TestNewMatchClient(t *testing.T) {
	c := NewClient(apiKey, EUROPE)
	_ = NewMatchClient(c)
}

func Test_matchClient_GetByPuuid(t *testing.T) {
	c := NewClient(apiKey, EUROPE)
	mc := NewMatchClient(c)
	matches, err := mc.GetMatchesByPuuid("5_QD37vWUa7Eq8jP6Cy-R18z60E9nRJlpkDZjqBGvtngedjANG6221udHyYnN2wCJCZV7CnlAqcnHQ")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if len(*matches) == 0 {
		t.Log("match history of player is empty, something went wrong")
		t.FailNow()
	}
}

func Test_matchClient_GetMatch(t *testing.T) {
	c := NewClient(apiKey, EUROPE)
	mc := NewMatchClient(c)
	matches, err := mc.GetMatch("EUW1_4770230362")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if matches.Metadata.MatchId == "" {
		t.Log("match id is empty, something went wrong")
		t.FailNow()
	}
}
