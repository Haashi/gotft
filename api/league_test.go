package api

import (
	"testing"
)

func TestNewLeagueClient(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	_ = NewLeagueClient(c, log)
}

func Test_leagueClient_GetMasterLeague(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetMasterLeague()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetMasterLeague_Fail(t *testing.T) {
	c := NewClient("wrongapikey", EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetMasterLeague()
	if err == nil || err.Code != ErrorUnauthorized {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorUnauthorized)
	}
}

func Test_leagueClient_GetGrandmasterLeague(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetGrandmasterLeague()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetGrandmasterLeague_Fail(t *testing.T) {
	c := NewClient("wrongapikey", EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetGrandmasterLeague()
	if err == nil || err.Code != ErrorUnauthorized {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorUnauthorized)
	}
}

func Test_leagueClient_GetChallengerLeague(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetChallengerLeague()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetChallengerLeague_Fail(t *testing.T) {
	c := NewClient("wrongapikey", EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetChallengerLeague()
	if err == nil || err.Code != ErrorUnauthorized {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorUnauthorized)
	}
}

func Test_leagueClient_GetBySummonerID(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetBySummonerID("cQqsUTIsR-TiXeV2LHALb5nx6tlma4UTavOj3u6KQseatVs")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetBySummonerID_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetBySummonerID("dQqsUTIsR-TiXeV2LHALb5nx6tlma4UTavOj3u6KQseatVs")
	if err == nil || err.Code != ErrorUnauthorized {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorUnauthorized)
	}
}

func Test_leagueClient_GetByTier(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetByTier(TierGold, DivI, 1)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetByTier_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetByTier(TierGold, "VI", 1)
	if err == nil || err.Code != ErrorUnauthorized {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorUnauthorized)
	}
}

func Test_leagueClient_GetById(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetById("dda7127a-bab9-3a9f-ad2f-7c798e3ce29a")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetById_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	lc := NewLeagueClient(c, log)
	_, err := lc.GetById("eda7127a-bab9-3a9f-ad2f-7c798e3ce29a")
	if err == nil || err.Code != ErrorNotFound {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorNotFound)
	}
}
