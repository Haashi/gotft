package match

import (
	"encoding/json"
	"fmt"

	"github.com/Haashi/gotft/api/client"
)

type MatchEndpoint struct {
	c *client.Client
}

func NewMatchEndpoint(c *client.Client) *MatchEndpoint {
	me := &MatchEndpoint{}
	me.c = c
	return me
}

func (me *MatchEndpoint) GetMatchesByPUUID(puuid string) *[]string {
	body, err := me.c.Get(fmt.Sprintf("/match/v1/matches/by-puuid/%s/ids", puuid), true)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	res := &[]string{}
	json.NewDecoder(body).Decode(res)
	return res
}
