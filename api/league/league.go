package league

import (
	"encoding/json"

	"github.com/Haashi/gotft/api/client"
)

type LeagueEndpoint struct {
	c *client.Client
}

func NewLeagueEndpoint(c *client.Client) *LeagueEndpoint {
	le := &LeagueEndpoint{}
	le.c = c
	return le
}

func (le *LeagueEndpoint) GetMasterLeague() *LeagueListDTO {
	body, err := le.c.Get("/league/v1/master", false)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	res := &LeagueListDTO{}
	json.NewDecoder(body).Decode(res)
	return res
}

func (le *LeagueEndpoint) GetGrandmasterLeague() *LeagueListDTO {
	body, err := le.c.Get("/league/v1/grandmaster", false)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	res := &LeagueListDTO{}
	json.NewDecoder(body).Decode(res)
	return res
}

func (le *LeagueEndpoint) GetChallengerLeague() *LeagueListDTO {
	body, err := le.c.Get("/league/v1/challenger", false)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	res := &LeagueListDTO{}
	json.NewDecoder(body).Decode(res)
	return res
}

type LeagueListDTO struct {
	LeagueId string
	Entries  []LeagueItemDTO
	Tier     string
	Name     string
	Queue    string
}

type LeagueItemDTO struct {
	FreshBlood   bool
	Wins         int
	SummonerName string
	MiniSeries   MiniSeriesDTO
	Inactive     bool
	Veteran      bool
	HotStreak    bool
	Rank         string
	LeaguePoints int
	Losses       int
	SummonerId   string
}

type MiniSeriesDTO struct {
	Losses   int
	Progress string
	Target   int
	Wins     int
}
