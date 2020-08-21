package api

import (
	"encoding/json"
	"fmt"
	"io"
)

type summonerClient struct {
	c *client
}

func NewSummonerClient(c *client) *summonerClient {
	sc := &summonerClient{}
	sc.c = c
	return sc
}

func (sc *summonerClient) GetByName(name string) (*Summoner, error) {
	body, err := sc.Get(fmt.Sprintf("/summoners/by-name/%s", name))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &Summoner{}
	json.NewDecoder(body).Decode(res)
	return res, nil
}

func (sc *summonerClient) GetByAccountId(accountId string) (*Summoner, error) {
	body, err := sc.Get(fmt.Sprintf("/summoners/by-account/%s", accountId))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &Summoner{}
	json.NewDecoder(body).Decode(res)
	return res, nil
}

func (sc *summonerClient) GetByPuuid(puuid string) (*Summoner, error) {
	body, err := sc.Get(fmt.Sprintf("/summoners/by-puuid/%s", puuid))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &Summoner{}
	json.NewDecoder(body).Decode(res)
	return res, nil
}

func (sc *summonerClient) GetById(id string) (*Summoner, error) {
	body, err := sc.Get(fmt.Sprintf("/summoners/%s", id))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &Summoner{}
	json.NewDecoder(body).Decode(res)
	return res, nil
}

func (sc *summonerClient) Get(url string) (io.ReadCloser, error) {
	return sc.c.Get(fmt.Sprintf("/summoner/v1%s", url))
}

type Summoner struct {
	AccountId     string `json:"accountId"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	Name          string `json:"name"`
	Id            string `json:"id"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
}
