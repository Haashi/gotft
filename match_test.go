package gotft

import (
	"testing"

	"github.com/haashi/gotft/internal"
)

func TestNewMatchClient(t *testing.T) {
	c := newClient(apiKey, EUROPE, testOpt)
	_ = newMatchClient(c, testOpt)
}

func Test_matchClient_GetByPuuid(t *testing.T) {
	c := newClient(apiKey, EUROPE, &Options{c: &internal.StringListClient{}, log: testOpt.log})
	mc := newMatchClient(c, testOpt)
	_, err := mc.GetMatchesByPuuid("puuid", 10)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_matchClient_GetByPuuid_Fail(t *testing.T) {
	c := newClient(apiKey, EUROPE, &Options{c: &internal.NotFoundClient{}, log: testOpt.log})
	mc := newMatchClient(c, testOpt)
	_, err := mc.GetMatchesByPuuid("puuid", 10)
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_matchClient_GetByPuuid_Decode(t *testing.T) {
	c := newClient(apiKey, EUROPE, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	mc := newMatchClient(c, testOpt)
	_, err := mc.GetMatchesByPuuid("puuid", 10)
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_matchClient_GetMatch(t *testing.T) {
	c := newClient(apiKey, EUROPE, &Options{c: &internal.MatchClient{}, log: testOpt.log})
	mc := newMatchClient(c, testOpt)
	_, err := mc.GetMatch("matchID")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_matchClient_GetMatch_Fail(t *testing.T) {
	c := newClient(apiKey, EUROPE, &Options{c: &internal.NotFoundClient{}, log: testOpt.log})
	mc := newMatchClient(c, testOpt)
	_, err := mc.GetMatch("matchID")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_matchClient_GetMatch_Decode(t *testing.T) {
	c := newClient(apiKey, EUROPE, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	mc := newMatchClient(c, testOpt)
	_, err := mc.GetMatch("matchID")
	if err == nil {
		t.Errorf("error missing")
	}
}
