package api

import (
	"encoding/json"
	"fmt"
	"io"
)

type leagueClient struct {
	c   *client
	log logger
}

func NewLeagueClient(c *client, log logger) *leagueClient {
	log.Debug("initializing league client")
	le := &leagueClient{}
	le.c = c
	le.log = log
	return le
}

func (le *leagueClient) GetMasterLeague() (*LeagueList, *Error) {
	le.log.Debug("getting master league")
	body, err := le.get("/master")
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		le.log.Errorf("error decoding master league : %s", errDec.Error())
		return nil, &Error{ErrorDecode, errDec.Error()}
	}
	return res, nil
}

func (le *leagueClient) GetGrandmasterLeague() (*LeagueList, *Error) {
	le.log.Debug("getting grandmaster league")
	body, err := le.get("/grandmaster")
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		le.log.Errorf("error decoding grandmaster league : %s", errDec.Error())
		return nil, &Error{ErrorDecode, errDec.Error()}
	}
	return res, nil
}

func (le *leagueClient) GetChallengerLeague() (*LeagueList, *Error) {
	le.log.Debug("getting challenger league")
	body, err := le.get("/challenger")
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		le.log.Errorf("error decoding challenger league : %s", errDec.Error())
		return nil, &Error{ErrorDecode, errDec.Error()}
	}
	return res, nil
}

func (le *leagueClient) GetBySummonerID(summonerId string) (*LeagueEntry, *Error) {
	le.log.Debugf("getting league of summonerid %s", summonerId)
	body, err := le.get(fmt.Sprintf("/entries/by-summoner/%s", summonerId))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &[]*LeagueEntry{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		le.log.Errorf("error decoding league of summonerid %s : %s", summonerId, errDec.Error())
		return nil, &Error{ErrorDecode, errDec.Error()}
	}
	return (*res)[0], nil
}

func (le *leagueClient) GetByTier(tier tier, division division, page int) (*[]*LeagueEntry, *Error) {
	le.log.Debugf("getting %s%s leagues (page%d)", tier, division, page)
	body, err := le.get(fmt.Sprintf("/entries/%s/%s?page=%d", tier, division, page))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &[]*LeagueEntry{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		le.log.Errorf("error decoding %s%s leagues (page%d) : %s", tier, division, page, errDec.Error())
		return nil, &Error{ErrorDecode, errDec.Error()}
	}
	return res, nil
}

func (le *leagueClient) GetById(id string) (*LeagueList, *Error) {
	le.log.Debugf("getting league %s", id)
	body, err := le.get(fmt.Sprintf("/leagues/%s", id))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		le.log.Errorf("error decoding league %s : %s", id, errDec.Error())
		return nil, &Error{ErrorDecode, errDec.Error()}
	}
	return res, nil
}

func (le *leagueClient) get(url string) (io.ReadCloser, *Error) {
	return le.c.Get(fmt.Sprintf("/league/v1%s", url))
}

type LeagueList struct {
	LeagueId string       `json:"leagueId" bson:"leagueId"`
	Entries  []LeagueItem `json:"entries" bson:"entries"`
	Tier     string       `json:"tier" bson:"tier"`
	Name     string       `json:"name" bson:"name"`
	Queue    string       `json:"queue" bson:"queue"`
}

type LeagueItem struct {
	FreshBlood   bool          `json:"freshBlood" bson:"freshBlood"`
	Wins         int           `json:"wins" bson:"wins"`
	SummonerName string        `json:"summonerName" bson:"summonerName"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries" bson:"miniSeries"`
	Inactive     bool          `json:"inactive" bson:"inactive"`
	Veteran      bool          `json:"veteran" bson:"veteran"`
	HotStreak    bool          `json:"hotStreak" bson:"hotStreak"`
	Rank         string        `json:"rank" bson:"rank"`
	LeaguePoints int           `json:"leaguePoints" bson:"leaguePoints"`
	Losses       int           `json:"losses" bson:"losses"`
	SummonerId   string        `json:"summonerId" bson:"summonerId"`
}

type LeagueEntry struct {
	LeagueId     string        `json:"leagueId" bson:"leagueId"`
	QueueType    string        `json:"queueType" bson:"queueType"`
	FreshBlood   bool          `json:"freshBlood" bson:"freshBlood"`
	Tier         string        `json:"tier" bson:"tier"`
	Wins         int           `json:"wins" bson:"wins"`
	SummonerName string        `json:"summonerName" bson:"summonerName"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries" bson:"miniSeries"`
	Inactive     bool          `json:"inactive" bson:"inactive"`
	Veteran      bool          `json:"veteran" bson:"veteran"`
	HotStreak    bool          `json:"hotStreak" bson:"hotStreak"`
	Rank         string        `json:"rank" bson:"rank"`
	LeaguePoints int           `json:"leaguePoints" bson:"leaguePoints"`
	Losses       int           `json:"losses" bson:"losses"`
	SummonerId   string        `json:"summonerId" bson:"summonerId"`
}

type MiniSeriesDTO struct {
	Losses   int    `json:"losses" bson:"losses"`
	Progress string `json:"progress" bson:"progress"`
	Target   int    `json:"target" bson:"target"`
	Wins     int    `json:"wins" bson:"wins"`
}

type division string

const (
	DivI   division = "I"
	DivII  division = "II"
	DivIII division = "III"
	DivIV  division = "IV"
)

type tier string

const (
	TierIron     tier = "IRON"
	TierBronze   tier = "BRONZE"
	TierSilver   tier = "SILVER"
	TierGold     tier = "GOLD"
	TierPlatinum tier = "PLATINUM"
	TierDiamond  tier = "DIAMOND"
)
