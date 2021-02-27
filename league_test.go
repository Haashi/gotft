package gotft

import (
	"testing"

	"github.com/haashi/gotft/internal"
)

func TestNewLeagueClient(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	_ = newLeagueClient(c, testOpt)
}

func Test_leagueClient_GetMasterLeague(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.LeagueListClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetMasterLeague()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetMasterLeague_Decode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetMasterLeague()
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetMasterLeague_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.NotFoundClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetMasterLeague()
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetGrandmasterLeague(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.LeagueListClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetGrandmasterLeague()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetGrandmasterLeague_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.NotFoundClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetGrandmasterLeague()
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetGrandmasterLeague_Decode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetGrandmasterLeague()
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetChallengerLeague(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.LeagueListClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetChallengerLeague()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetChallengerLeague_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadRequestClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetChallengerLeague()
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetChallengerLeague_Decode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetChallengerLeague()
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetBySummonerID(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.LeagueEntriesClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetBySummonerID("summonerID")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetBySummonerID_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.UnauthorizedClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetBySummonerID("summonerID")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetBySummonerID_Decode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetBySummonerID("summonerID")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetByTier(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.LeagueEntriesClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetByTier(TierGold, DivI, 1)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetByTier_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.TimeoutClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetByTier(TierGold, "VI", 1)
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetByTier_Decode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetByTier(TierGold, "VI", 1)
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetById(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.LeagueListClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetById("ID")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_leagueClient_GetById_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.ForbiddenClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetById("ID")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_leagueClient_GetById_Decode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	lc := newLeagueClient(c, testOpt)
	_, err := lc.GetById("ID")
	if err == nil {
		t.Errorf("error missing")
	}
}
