package gotft

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/haashi/gotft/internal"
)

type matchClient struct {
	c   *apiclient
	log internal.Logger
}

func newMatchClient(c *apiclient, opt *Options) *matchClient {
	opt.log.Debug("initializing match client")
	mc := &matchClient{}
	mc.c = c
	mc.log = opt.log
	return mc
}

func (mc *matchClient) GetMatchesByPuuid(puuid string, count int) (*[]string, error) {
	mc.log.Debugf("getting matches list(%d) of puuid %s", count, puuid)
	body, err := mc.get(fmt.Sprintf("/matches/by-puuid/%s/ids?count=%d", puuid, count))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &[]string{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		mc.log.Errorf("error decoding matches list(%d) of puuid %s : %s", count, puuid, errDec.Error())
		return nil, ErrorDecode{fmt.Sprintf("matches list(%d) of puuid %s", count, puuid), errDec.Error()}
	}
	return res, nil
}

func (mc *matchClient) GetMatch(id string) (*Match, error) {
	mc.log.Debugf("getting match %s", id)
	body, err := mc.get(fmt.Sprintf("/matches/%s", id))
	if err != nil {
		mc.log.Errorf("error getting match %s : %s", id, err.Error())
		return nil, err
	}
	defer body.Close()
	res := &Match{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		mc.log.Errorf("error decoding match %s : %s", id, errDec.Error())
		return nil, ErrorDecode{fmt.Sprintf("match %s", id), errDec.Error()}
	}
	return res, nil
}

func (mc *matchClient) get(url string) (io.ReadCloser, error) {
	return mc.c.get(fmt.Sprintf("/match/v1%s", url))
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
