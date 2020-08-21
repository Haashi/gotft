package api

import (
	"testing"
)

func TestNewLeagueClient(t *testing.T) {
	c := NewClient(apiKey, EUW)
	_ = NewLeagueClient(c)
}

func Test_leagueClient_GetMasterLeague(t *testing.T) {
	c := NewClient(apiKey, EUW)
	lc := NewLeagueClient(c)
	master, err := lc.GetMasterLeague()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if master.LeagueId == "" {
		t.Log("master leagueID is empty, something went wrong")
		t.FailNow()
	}
}

func Test_leagueClient_GetGrandmasterLeague(t *testing.T) {
	c := NewClient(apiKey, EUW)
	lc := NewLeagueClient(c)
	grandmaster, err := lc.GetGrandmasterLeague()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if grandmaster.LeagueId == "" {
		t.Log("grandmaster leagueID is empty, something went wrong")
		t.FailNow()
	}
}

func Test_leagueClient_GetChallengerLeague(t *testing.T) {
	c := NewClient(apiKey, EUW)
	lc := NewLeagueClient(c)
	challenger, err := lc.GetChallengerLeague()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if challenger.LeagueId == "" {
		t.Log("challenger leagueID is empty, something went wrong")
		t.FailNow()
	}
}
