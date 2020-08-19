package summoner

import (
	"encoding/json"
	"fmt"

	"github.com/Haashi/gotft/api/client"
)

type SummonerEndpoint struct {
	c *client.Client
}

func NewSummonerEndpoint(c *client.Client) *SummonerEndpoint {
	se := &SummonerEndpoint{}
	se.c = c
	return se
}

func (se *SummonerEndpoint) GetByName(name string) *SummonerDTO {
	body, err := se.c.Get(fmt.Sprintf("/summoner/v1/summoners/by-name/%s", name), false)
	if err != nil {
		panic(err)
	}
	defer body.Close()
	res := &SummonerDTO{}
	json.NewDecoder(body).Decode(res)
	return res
}

type SummonerDTO struct {
	AccountId     string
	ProfileIconId int
	RevisionDate  int64
	Name          string
	Id            string
	Puuid         string
	SummonerLevel int
}
