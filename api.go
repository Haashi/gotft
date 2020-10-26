package gotft

import "github.com/haashi/gotft/internal"

type region string

const (
	BR       region = "br1"
	EUN      region = "eun1"
	EUW      region = "euw1"
	JP       region = "jp1"
	KR       region = "kr"
	LAN      region = "la1"
	LAS      region = "la2"
	NA       region = "na1"
	OCE      region = "oc1"
	RU       region = "ru"
	TR       region = "tr1"
	EUROPE   region = "europe"
	ASIA     region = "asia"
	AMERICAS region = "americas"
)

var regionToContinent = map[region]region{
	BR:  AMERICAS,
	EUN: EUROPE,
	EUW: EUROPE,
	JP:  ASIA,
	KR:  ASIA,
	LAN: AMERICAS,
	LAS: AMERICAS,
	NA:  AMERICAS,
	OCE: ASIA,
	RU:  EUROPE,
	TR:  EUROPE,
}

type api struct {
	League   *leagueClient
	Match    *matchClient
	Summoner *summonerClient
}

func newAPI(apiKey string, region region, options ...Option) *api {
	fields := map[string]interface{}{
		"region": region,
	}
	opt := &Options{
		log: internal.NewDefaultLogger(fields),
		c:   internal.NewDefaultClient(),
	}
	for _, option := range options {
		option(opt)
	}
	api := &api{}
	baseClient := newClient(apiKey, region, opt)
	continentClient := newClient(apiKey, regionToContinent[region], opt)
	api.League = newLeagueClient(baseClient, opt)
	api.Match = newMatchClient(continentClient, opt)
	api.Summoner = newSummonerClient(baseClient, opt)
	return api
}
