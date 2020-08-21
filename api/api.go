package api

const (
	BR       = "br1"
	EUN      = "eun1"
	EUW      = "euw1"
	JP       = "jp1"
	KR       = "kr"
	LAN      = "la1"
	LAS      = "la2"
	NA       = "na1"
	OCE      = "oc1"
	RU       = "ru"
	TR       = "tr1"
	EUROPE   = "europe"
	ASIA     = "asia"
	AMERICAS = "americas"
)

var regionToContinent = map[string]string{
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
}

type Key struct {
	Value string `json:"value"`
	Prod  bool   `json:"prod"`
}

func NewAPI(apikey string, region string) *API {
	baseClient := NewClient(apikey, region)
	continentClient := NewClient(apikey, regionToContinent[region])
	api := &API{}
	api.League = NewLeagueClient(baseClient)
	api.Match = NewMatchClient(continentClient)
	api.Summoner = NewSummonerClient(baseClient)
	return api
}
