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

func (mc *matchClient) GetMatchesByPuuid(puuid string, count int) (*[]string, error) {
	body, err := mc.get(fmt.Sprintf("/matches/by-puuid/%s/ids?count=%d", puuid, count))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &[]string{}
	json.NewDecoder(body).Decode(res)
	return res, nil
}

func (mc *matchClient) GetMatch(id string) (*Match, error) {
	body, err := mc.get(fmt.Sprintf("/matches/%s", id))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &Match{}
	json.NewDecoder(body).Decode(res)
	return res, nil
}

func (mc *matchClient) get(url string) (io.ReadCloser, error) {
	return mc.c.Get(fmt.Sprintf("/match/v1%s", url))
}

type Match struct {
	Metadata Metadata `json:"metadata" bson:"metadata"`
	Info     Info     `json:"info" bson:"info"`
}

type Metadata struct {
	DataVersion  string   `json:"data_version" bson:"data_version"`
	MatchId      string   `json:"match_id" bson:"match_id"`
	Participants []string `json:"participants" bson:"participants"`
}

type Info struct {
	GameDatetime  int64         `json:"game_datetime" bson:"game_datetime"`
	GameLength    float64       `json:"game_length" bson:"game_length"`
	GameVariation string        `json:"game_variation" bson:"game_variation"`
	GameVersion   string        `json:"game_version" bson:"game_version"`
	Participants  []Participant `json:"participants" bson:"participants"`
	QueueId       int           `json:"queue_id" bson:"queue_id"`
	TftSetNumber  int           `json:"tft_set_number" bson:"tft_set_number"`
}

type Companion struct {
	ContentId string `json:"content_ID" bson:"content_ID"`
	SkinID    int    `json:"skin_ID" bson:"skin_ID"`
	Species   string `json:"species" bson:"species"`
}

type Participant struct {
	Companion           Companion `json:"companion" bson:"companion"`
	GoldLeft            int       `json:"gold_left" bson:"gold_left"`
	LastRound           int       `json:"last_round" bson:"last_round"`
	Level               int       `json:"level" bson:"level"`
	Placement           int       `json:"placement" bson:"placement"`
	PlayersEliminated   int       `json:"players_eliminated" bson:"players_eliminated"`
	Puuid               string    `json:"puuid" bson:"puuid"`
	TimeEliminated      float64   `json:"time_eliminated" bson:"time_eliminated"`
	TotalDamageToPlayer int       `json:"total_damage_to_players" bson:"total_damage_to_players"`
	Traits              []Trait   `json:"traits" bson:"traits"`
	Units               []Unit    `json:"units" bson:"units"`
}

type Trait struct {
	Name        string `json:"name" bson:"name"`
	NumUnits    int    `json:"num_units" bson:"num_units"`
	TierCurrent int    `json:"tier_current" bson:"tier_current"`
	TierTotal   int    `json:"tier_total" bson:"tier_total"`
	Style       int    `json:"style" bson:"style"`
}

type Unit struct {
	Items       []int  `json:"items" bson:"items"`
	CharacterId string `json:"character_id" bson:"character_id"`
	Name        string `json:"name" bson:"name"`
	Rarity      int    `json:"rarity" bson:"rarity"`
	Tier        int    `json:"tier" bson:"tier"`
}
