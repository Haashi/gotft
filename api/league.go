package api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/haashi/gotft/api/internal"
)

type leagueClient struct {
	c   *apiclient
	log internal.Logger
}

func newLeagueClient(c *apiclient, opt *Options) *leagueClient {
	opt.log.Debug("initializing league client")
	le := &leagueClient{}
	le.c = c
	le.log = opt.log
	return le
}

func (le *leagueClient) GetMasterLeague() (*LeagueList, error) {
	le.log.Debug("getting master league")
	body, err := le.get("/master")
	if err != nil {
		le.log.Errorf("error getting master league : %s", err.Error())
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		err := ErrorDecode{"master league", errDec.Error()}
		le.log.Errorf(err.Error())
		return nil, err
	}
	return res, nil
}

func (le *leagueClient) GetGrandmasterLeague() (*LeagueList, error) {
	le.log.Debug("getting grandmaster league")
	body, err := le.get("/grandmaster")
	if err != nil {
		le.log.Errorf("error getting grandmaster league : %s", err.Error())
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		err := ErrorDecode{"grandmaster league", errDec.Error()}
		le.log.Errorf(err.Error())
		return nil, err
	}
	return res, nil
}

func (le *leagueClient) GetChallengerLeague() (*LeagueList, error) {
	le.log.Debug("getting challenger league")
	body, err := le.get("/challenger")
	if err != nil {
		le.log.Errorf("error getting challenger league : %s", err.Error())
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		err := ErrorDecode{"challenger league", errDec.Error()}
		le.log.Errorf(err.Error())
		return nil, err
	}
	return res, nil
}

func (le *leagueClient) GetBySummonerID(summonerId string) (*LeagueEntry, error) {
	le.log.Debugf("getting league of summonerid %s", summonerId)
	body, err := le.get(fmt.Sprintf("/entries/by-summoner/%s", summonerId))
	if err != nil {
		le.log.Errorf("error getting league league of summonerid %s : %s", summonerId, err.Error())
		return nil, err
	}
	defer body.Close()
	res := &[]*LeagueEntry{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		err := ErrorDecode{fmt.Sprintf("league of summonerid %s", summonerId), errDec.Error()}
		le.log.Errorf(err.Error())
		return nil, err
	}
	return (*res)[0], nil
}

func (le *leagueClient) GetByTier(tier tier, division division, page int) (*[]*LeagueEntry, error) {
	le.log.Debugf("getting %s%s leagues (page%d)", tier, division, page)
	body, err := le.get(fmt.Sprintf("/entries/%s/%s?page=%d", tier, division, page))
	if err != nil {
		le.log.Errorf("error getting %s%s leagues (page%d) : %s", tier, division, page, err.Error())
		return nil, err
	}
	defer body.Close()
	res := &[]*LeagueEntry{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		err := ErrorDecode{fmt.Sprintf("%s%s leagues (page%d)", tier, division, page), errDec.Error()}
		le.log.Errorf(err.Error())
		return nil, err
	}
	return res, nil
}

func (le *leagueClient) GetById(id string) (*LeagueList, error) {
	le.log.Debugf("getting league %s", id)
	body, err := le.get(fmt.Sprintf("/leagues/%s", id))
	if err != nil {
		le.log.Errorf("error getting league %s : %s", id, err.Error())
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		err := ErrorDecode{fmt.Sprintf("league %s", id), errDec.Error()}
		le.log.Errorf(err.Error())
		return nil, err
	}
	return res, nil
}

func (le *leagueClient) get(url string) (io.ReadCloser, error) {
	return le.c.get(fmt.Sprintf("/league/v1%s", url))
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
