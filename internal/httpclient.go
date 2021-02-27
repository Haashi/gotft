package internal

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func NewDefaultClient() HttpClient {
	return &http.Client{}
}

type TimeoutClient struct {
}

func (c TimeoutClient) Do(r *http.Request) (*http.Response, error) {
	return nil, http.ErrHandlerTimeout
}

type NotFoundClient struct {
}

func (c NotFoundClient) Do(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusNotFound}, nil
}

type UnauthorizedClient struct {
}

func (c UnauthorizedClient) Do(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusUnauthorized}, nil
}

type ForbiddenClient struct {
}

func (c ForbiddenClient) Do(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusForbidden}, nil
}

type TooManyRequestsClient struct {
	count int
}

func (c *TooManyRequestsClient) Do(r *http.Request) (*http.Response, error) {
	if c.count == 0 {
		c.count += 1
		headers := http.Header{}
		headers.Add("Retry-After", "1")
		return &http.Response{StatusCode: http.StatusTooManyRequests, Header: headers}, nil
	} else {
		return &http.Response{StatusCode: http.StatusOK}, nil
	}
}

type BadRequestClient struct {
}

func (c *BadRequestClient) Do(r *http.Request) (*http.Response, error) {
	ll := `{
		"status_code": 200,
		"message": "This is a message from the mock api server."
	}`
	return &http.Response{StatusCode: http.StatusBadRequest, Body: ioutil.NopCloser(strings.NewReader(ll))}, nil
}

type LeagueListClient struct {
}

func (c LeagueListClient) Do(r *http.Request) (*http.Response, error) {
	ll := `{
					"tier": "MASTER",
					"leagueId": "33829522-2e51-3f34-904e-4f5c14a2c2c8",
					"queue": "RANKED_TFT",
					"name": "Maokai's Captains",
					"entries":
						[
							{
									"summonerId": "summonerId",
									"summonerName": "SummonerName",
									"leaguePoints": 0,
									"rank": "I",
									"wins": 10,
									"losses": 10,
									"veteran": false,
									"inactive": false,
									"freshBlood": false,
									"hotStreak": false
							}
						]
					}`
	return &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(ll))}, nil
}

type LeagueEntriesClient struct {
}

func (c LeagueEntriesClient) Do(r *http.Request) (*http.Response, error) {
	le := `[
					{
						"leagueId": "leagueID",
						"queueType": "RANKED_TFT",
						"tier": "MASTER",
						"rank": "I",
						"summonerId": "summonerID",
						"summonerName": "summonerName",
						"leaguePoints": 0,
						"wins": 10,
						"losses": 10,
						"veteran": false,
						"inactive": false,
						"freshBlood": false,
						"hotStreak": false
					}
				]`
	return &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(le))}, nil
}

type StringListClient struct {
}

func (c StringListClient) Do(r *http.Request) (*http.Response, error) {
	sl := `[
					"match1",
					"match2",
					"match3",
					"match4",
					"match5",
					"match6",
					"match7",
					"match8",
					"match9",
					"match10"
				]`
	return &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(sl))}, nil
}

type MatchClient struct {
}

func (c MatchClient) Do(r *http.Request) (*http.Response, error) {
	m := `{
		"metadata": {
			"data_version": "5",
			"match_id": "matchid",
			"participants": [
				"puuid"
			]
		},
		"info": {
			"game_datetime": 1607975224052,
			"game_length": 2154.652099609375,
			"game_version": "Version",
			"participants": [
				{
					"companion": {
						"content_ID": "a4a35782-fbae-45ee-bc0d-90c26269d1a9",
						"skin_ID": 18,
						"species": "PetChoncc"
					},
					"gold_left": 0,
					"last_round": 33,
					"level": 8,
					"placement": 5,
					"players_eliminated": 0,
					"puuid": "puuid",
					"time_eliminated": 1905.9017333984375,
					"total_damage_to_players": 63,
					"traits": [
						{
							"name": "Cultist",
							"num_units": 1,
							"style": 0,
							"tier_current": 0,
							"tier_total": 3
						}
					],
					"units": [
						{
							"character_id": "characterid",
							"chosen": "Cultist",
							"items": [
								17
							],
							"name": "",
							"rarity": 0,
							"tier": 2
						}
					]
				}
			],
			"queue_id": 1100,
			"tft_set_number": 4
		}
	}`
	return &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(m))}, nil
}

type SummonerClient struct {
}

func (c SummonerClient) Do(r *http.Request) (*http.Response, error) {
	s := `{
					"id": "id",
					"accountId": "accountid",
					"puuid": "puuid",
					"name": "name",
					"profileIconId": 1,
					"revisionDate": 1608205772000,
					"summonerLevel": 100
				}`
	return &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(s))}, nil
}

type BadRequestBadDecodeClient struct {
}

func (c BadRequestBadDecodeClient) Do(r *http.Request) (*http.Response, error) {
	s := `---
	status_code: 200
	message: This is a message from the mock api server.	
	`
	return &http.Response{StatusCode: http.StatusBadRequest, Body: ioutil.NopCloser(strings.NewReader(s))}, nil
}

type BadDecodeClient struct {
}

func (c BadDecodeClient) Do(r *http.Request) (*http.Response, error) {
	s := `---
	`
	return &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(s))}, nil
}
