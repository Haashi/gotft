package api

import (
	"testing"
)

func TestNewSummonerClient(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	_ = NewSummonerClient(c, log)
}

func Test_summonerClient_GetByName(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	sc := NewSummonerClient(c, log)
	_, err := sc.GetByName("Haashi")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_summonerClient_GetByName_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	sc := NewSummonerClient(c, log)
	_, err := sc.GetByName("Haashiiiiiiiiiiiii")
	if err == nil || err.Code != ErrorNotFound {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorNotFound)
	}
}

func Test_summonerClient_GetByAccountId(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	sc := NewSummonerClient(c, log)
	_, err := sc.GetByAccountId("KGtzAbSB_J3H2S2wYoDN51BRpJLHqyBx4vV7bJ8Yu14v8g")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_summonerClient_GetByAccountId_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	sc := NewSummonerClient(c, log)
	_, err := sc.GetByAccountId("BGtzAbSB_J3H2S2wYoDN51BRpJLHqyBx4vV7bJ8Yu14v8g")
	if err == nil || err.Code != ErrorUnauthorized {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorUnauthorized)
	}
}

func Test_summonerClient_GetByPuuid(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	sc := NewSummonerClient(c, log)
	_, err := sc.GetByPuuid("5_QD37vWUa7Eq8jP6Cy-R18z60E9nRJlpkDZjqBGvtngedjANG6221udHyYnN2wCJCZV7CnlAqcnHQ")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_summonerClient_GetByPuuid_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	sc := NewSummonerClient(c, log)
	_, err := sc.GetByPuuid("6_QD37vWUa7Eq8jP6Cy-R18z60E9nRJlpkDZjqBGvtngedjANG6221udHyYnN2wCJCZV7CnlAqcnHQ")
	if err == nil || err.Code != ErrorUnauthorized {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorUnauthorized)
	}
}

func Test_summonerClient_GetById(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	sc := NewSummonerClient(c, log)
	_, err := sc.GetById("cQqsUTIsR-TiXeV2LHALb5nx6tlma4UTavOj3u6KQseatVs")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_summonerClient_GetById_Fail(t *testing.T) {
	c := NewClient(apiKey, EUW, log)
	sc := NewSummonerClient(c, log)
	_, err := sc.GetById("dQqsUTIsR-TiXeV2LHALb5nx6tlma4UTavOj3u6KQseatVs")
	if err == nil || err.Code != ErrorUnauthorized {
		t.Errorf("error code is %d, expected %d", err.Code, ErrorUnauthorized)
	}
}
