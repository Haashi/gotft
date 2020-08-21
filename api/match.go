package api

import (
	"encoding/json"
	"fmt"
	"io"
)

type matchClient struct {
	c *client
}

func NewMatchClient(c *client) *matchClient {
	mc := &matchClient{}
	mc.c = c
	return mc
}

func (mc *matchClient) GetMatchesByPuuid(puuid string) (*[]string, error) {
	body, err := mc.Get(fmt.Sprintf("/matches/by-puuid/%s/ids", puuid))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &[]string{}
	json.NewDecoder(body).Decode(res)
	return res, nil
}

func (mc *matchClient) GetMatch(id string) (*Match, error) {
	body, err := mc.Get(fmt.Sprintf("/matches/%s", id))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &Match{}
	json.NewDecoder(body).Decode(res)
	return res, nil
}

func (mc *matchClient) Get(url string) (io.ReadCloser, error) {
	return mc.c.Get(fmt.Sprintf("/match/v1%s", url))
}

type Match struct {
	Metadata Metadata `json:"metadata"`
	Info     Info     `json:"info"`
}

type Metadata struct {
	DataVersion  string   `json:"data_version"`
	MatchId      string   `json:"match_id"`
	Participants []string `json:"participants"`
}

type Info struct {
	GameDatetime  int64         `json:"game_datetime"`
	GameLength    float64       `json:"game_length"`
	GameVariation string        `json:"game_variation"`
	GameVersion   string        `json:"game_version"`
	Participants  []Participant `json:"participants"`
	QueueId       int           `json:"queue_id"`
	TftSetNumber  int           `json:"tft_set_number"`
}

type Companion struct {
	ContentId string `json:"content_ID"`
	SkinID    string `json:"skin_ID"`
	Species   string `json:"species"`
}

type Participant struct {
	Companion           Companion `json:"companion"`
	GoldLeft            int       `json:"gold_left"`
	LastRound           int       `json:"last_round"`
	Level               int       `json:"level"`
	Placement           int       `json:"placement"`
	PlayersEliminated   int       `json:"players_eliminated"`
	Puuid               string    `json:"puuid"`
	TimeEliminated      float64   `json:"time_eliminated"`
	TotalDamageToPlayer int       `json:"total_damage_to_players"`
	Traits              []Trait   `json:"traits"`
	Units               []Unit    `json:"units"`
}

type Trait struct {
	Name        string `json:"name"`
	NumUnits    int    `json:"num_units"`
	TierCurrent int    `json:"tier_current"`
	TierTotal   int    `json:"tier_total"`
}

type Unit struct {
	Items       []int  `json:"items"`
	CharacterId string `json:"character_id"`
	Name        string `json:"name"`
	Rarity      int    `json:"rarity"`
	Tier        int    `json:"tier"`
}
