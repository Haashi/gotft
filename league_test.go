package gotft

import (
	"testing"
)

func TestNewLeagueClient(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	_ = newLeagueClient(c, testOpt)
}

func Test_leagueClient_GetMasterLeague(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetMasterLeague()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetMasterLeague_Fail(t *testing.T) {
	c := newClient("RGAPI-a051aef9-0b01-4bc1-b7e0-fbf23eaafd1d", EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetMasterLeague()
	if _, ok := err.(ErrorForbidden); !ok {
		t.Errorf("error missing or error is not a forbidden error")
	}
}

func Test_leagueClient_GetGrandmasterLeague(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetGrandmasterLeague()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetGrandmasterLeague_Fail(t *testing.T) {
	c := newClient("wrongapikey", EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetGrandmasterLeague()
	if _, ok := err.(ErrorForbidden); !ok {
		t.Errorf("error missing or error is not a forbidden error")
	}
}

func Test_leagueClient_GetChallengerLeague(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetChallengerLeague()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetChallengerLeague_Fail(t *testing.T) {
	c := newClient("wrongapikey", EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetChallengerLeague()
	if _, ok := err.(ErrorForbidden); !ok {
		t.Errorf("error missing or error is not a forbidden error")
	}
}

func Test_leagueClient_GetBySummonerID(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetBySummonerID("cQqsUTIsR-TiXeV2LHALb5nx6tlma4UTavOj3u6KQseatVs")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetBySummonerID_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetBySummonerID("dQqsUTIsR-TiXeV2LHALb5nx6tlma4UTavOj3u6KQseatVs")
	if _, ok := err.(ErrorBadRequest); !ok {
		t.Errorf("error missing or error is not a badrequest error")
	}
}

func Test_leagueClient_GetByTier(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetByTier(TierGold, DivI, 1)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetByTier_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetByTier(TierGold, "VI", 1)
	if _, ok := err.(ErrorBadRequest); !ok {
		t.Errorf("error missing or error is not a badrequest error")
	}
}

func Test_leagueClient_GetById(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetById("dda7127a-bab9-3a9f-ad2f-7c798e3ce29a")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetById_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetById("eda7127a-bab9-3a9f-ad2f-7c798e3ce29a")
	if _, ok := err.(ErrorNotFound); !ok {
		t.Errorf("error missing or error is not a notfound error")
	}
}
