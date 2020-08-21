package api

import (
	"encoding/json"
	"fmt"
	"io"
)

type leagueClient struct {
	c *client
}

func NewLeagueClient(c *client) *leagueClient {
	le := &leagueClient{}
	le.c = c
	return le
}

func (le *leagueClient) GetMasterLeague() (*LeagueList, error) {
	body, err := le.Get("/master")
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	err = json.NewDecoder(body).Decode(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (le *leagueClient) GetGrandmasterLeague() (*LeagueList, error) {
	body, err := le.Get("/grandmaster")
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	err = json.NewDecoder(body).Decode(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (le *leagueClient) GetChallengerLeague() (*LeagueList, error) {
	body, err := le.Get("/challenger")
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &LeagueList{}
	json.NewDecoder(body).Decode(res)
	return res, nil
}

func (le *leagueClient) Get(url string) (io.ReadCloser, error) {
	return le.c.Get(fmt.Sprintf("/league/v1%s", url))
}

type LeagueList struct {
	LeagueId string       `json:"leagueId"`
	Entries  []LeagueItem `json:"entries"`
	Tier     string       `json:"tier"`
	Name     string       `json:"name"`
	Queue    string       `json:"queue"`
}

type LeagueItem struct {
	FreshBlood   bool          `json:"freshBlood"`
	Wins         int           `json:"wins"`
	SummonerName string        `json:"summonerName"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries"`
	Inactive     bool          `json:"inactive"`
	Veteran      bool          `json:"veteran"`
	HotStreak    bool          `json:"hotStreak"`
	Rank         string        `json:"rank"`
	LeaguePoints int           `json:"leaguePoints"`
	Losses       int           `json:"losses"`
	SummonerId   string        `json:"summonerId"`
}

type MiniSeriesDTO struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}
