package api

import (
	"testing"

	"github.com/haashi/gotft/api/internal"
)

func TestNewSummonerClient(t *testing.T) {
	c := newClient(apiKey, EUW, testOpt)
	_ = newSummonerClient(c, testOpt)
}

func Test_summonerClient_GetByName(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.SummonerClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetByName("name")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_summonerClient_GetByName_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.NotFoundClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetByName("name")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_summonerClient_GetByName_Decode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetByName("name")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_summonerClient_GetByAccountId(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.SummonerClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetByAccountId("accountid")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_summonerClient_GetByAccountId_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.NotFoundClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetByAccountId("accountid")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_summonerClient_GetByAccountId_Decode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetByAccountId("accountid")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_summonerClient_GetByPuuid(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.SummonerClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetByPuuid("puuid")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_summonerClient_GetByPuuid_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.NotFoundClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetByPuuid("puuid")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_summonerClient_GetByPuuid_Decode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetByPuuid("puuid")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_summonerClient_GetById(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.SummonerClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetById("id")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_summonerClient_GetById_Fail(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.NotFoundClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetById("id")
	if err == nil {
		t.Errorf("error missing")
	}
}

func Test_summonerClient_GetById_Decode(t *testing.T) {
	c := newClient(apiKey, EUW, &Options{c: &internal.BadDecodeClient{}, log: testOpt.log})
	sc := newSummonerClient(c, testOpt)
	_, err := sc.GetById("id")
	if err == nil {
		t.Errorf("error missing")
	}
}
