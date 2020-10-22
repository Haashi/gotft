package api

import (
	"github.com/sirupsen/logrus"
)

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

type API struct {
	League   *leagueClient
	Match    *matchClient
	Summoner *summonerClient
	log      logger
}

type Key struct {
	Value string `json:"value"`
	Prod  bool   `json:"prod"`
}

func NewAPI(apikey string, region region, options ...Option) *API {
	defaultLogger := logrus.New()
	defaultLogger.SetLevel(logrus.ErrorLevel)
	opt := &Options{
		log: defaultLogger,
	}
	for _, option := range options {
		option(opt)
	}
	api := &API{}
	api.log = opt.log
	baseClient := NewClient(apikey, region, api.log)
	continentClient := NewClient(apikey, regionToContinent[region], api.log)
	api.League = NewLeagueClient(baseClient, api.log)
	api.Match = NewMatchClient(continentClient, api.log)
	api.Summoner = NewSummonerClient(baseClient, api.log)
	return api
}
