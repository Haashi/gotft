package api

import (
	"github.com/Haashi/gotft/api/client"
	"github.com/Haashi/gotft/api/league"
	"github.com/Haashi/gotft/api/match"
	"github.com/Haashi/gotft/api/summoner"
)

type API struct {
	League   *league.LeagueEndpoint
	Match    *match.MatchEndpoint
	Summoner *summoner.SummonerEndpoint
}

func NewAPI(apikey string, region string, prodKey bool) *API {
	c := client.NewClient(apikey, region, prodKey)
	api := &API{}
	api.League = league.NewLeagueEndpoint(c)
	api.Match = match.NewMatchEndpoint(c)
	api.Summoner = summoner.NewSummonerEndpoint(c)
	return api
}
