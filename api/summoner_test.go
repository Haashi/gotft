package api

import (
	"testing"
)

func TestNewSummonerClient(t *testing.T) {
	c := NewClient(apiKey, EUW)
	_ = NewSummonerClient(c)
}

func Test_summonerClient_GetByName(t *testing.T) {
	c := NewClient(apiKey, EUW)
	sc := NewSummonerClient(c)
	summoner, err := sc.GetByName("Haashi")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if summoner.AccountId == "" {
		t.Log("summoner account id is empty, something went wrong")
		t.FailNow()
	}
}

func Test_summonerClient_GetByName_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW)
	sc := NewSummonerClient(c)
	_, err := sc.GetByName("Haashiiiiiiiiiiiii")
	if err == nil {
		t.Log(err)
		t.FailNow()
	}
}

func Test_summonerClient_GetByAccountId(t *testing.T) {
	c := NewClient(apiKey, EUW)
	sc := NewSummonerClient(c)
	summoner, err := sc.GetByAccountId("KGtzAbSB_J3H2S2wYoDN51BRpJLHqyBx4vV7bJ8Yu14v8g")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if summoner.AccountId == "" {
		t.Log("summoner account id is empty, something went wrong")
		t.FailNow()
	}
}

func Test_summonerClient_GetByAccountId_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW)
	sc := NewSummonerClient(c)
	_, err := sc.GetByAccountId("BGtzAbSB_J3H2S2wYoDN51BRpJLHqyBx4vV7bJ8Yu14v8g")
	if err == nil {
		t.Log(err)
		t.FailNow()
	}
}

func Test_summonerClient_GetByPuuid(t *testing.T) {
	c := NewClient(apiKey, EUW)
	sc := NewSummonerClient(c)
	summoner, err := sc.GetByPuuid("5_QD37vWUa7Eq8jP6Cy-R18z60E9nRJlpkDZjqBGvtngedjANG6221udHyYnN2wCJCZV7CnlAqcnHQ")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if summoner.AccountId == "" {
		t.Log("summoner account id is empty, something went wrong")
		t.FailNow()
	}
}

func Test_summonerClient_GetByPuuid_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW)
	sc := NewSummonerClient(c)
	_, err := sc.GetByPuuid("6_QD37vWUa7Eq8jP6Cy-R18z60E9nRJlpkDZjqBGvtngedjANG6221udHyYnN2wCJCZV7CnlAqcnHQ")
	if err == nil {
		t.Log(err)
		t.FailNow()
	}
}

func Test_summonerClient_GetById(t *testing.T) {
	c := NewClient(apiKey, EUW)
	sc := NewSummonerClient(c)
	summoner, err := sc.GetById("cQqsUTIsR-TiXeV2LHALb5nx6tlma4UTavOj3u6KQseatVs")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if summoner.AccountId == "" {
		t.Log("summoner account id is empty, something went wrong")
		t.FailNow()
	}
}

func Test_summonerClient_GetById_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW)
	sc := NewSummonerClient(c)
	_, err := sc.GetById("dQqsUTIsR-TiXeV2LHALb5nx6tlma4UTavOj3u6KQseatVs")
	if err == nil {
		t.Log(err)
		t.FailNow()
	}
}
