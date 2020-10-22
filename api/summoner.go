package api

import (
	"encoding/json"
	"fmt"
	"io"
)

type summonerClient struct {
	c   *client
	log logger
}

func NewSummonerClient(c *client, log logger) *summonerClient {
	log.Debug("initializing summoner client")
	sc := &summonerClient{}
	sc.c = c
	sc.log = log
	return sc
}

func (sc *summonerClient) GetByName(name string) (*Summoner, *Error) {
	sc.log.Debugf("getting summoner by name %s", name)
	body, err := sc.get(fmt.Sprintf("/summoners/by-name/%s", name))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &Summoner{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		sc.log.Errorf("error decoding summoner by name %s : %s", name, errDec.Error())
		return nil, &Error{ErrorDecode, errDec.Error()}
	}
	return res, nil
}

func (sc *summonerClient) GetByAccountId(accountId string) (*Summoner, *Error) {
	sc.log.Debugf("getting summoner by accountid %s", accountId)
	body, err := sc.get(fmt.Sprintf("/summoners/by-account/%s", accountId))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &Summoner{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		sc.log.Errorf("error decoding summoner by accountid %s : %s", accountId, errDec.Error())
		return nil, &Error{ErrorDecode, errDec.Error()}
	}
	return res, nil
}

func (sc *summonerClient) GetByPuuid(puuid string) (*Summoner, *Error) {
	sc.log.Debugf("getting summoner by puuid %s", puuid)
	body, err := sc.get(fmt.Sprintf("/summoners/by-puuid/%s", puuid))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &Summoner{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		sc.log.Errorf("error decoding summoner by puuid %s : %s", puuid, errDec.Error())
		return nil, &Error{ErrorDecode, errDec.Error()}
	}
	return res, nil
}

func (sc *summonerClient) GetById(id string) (*Summoner, *Error) {
	sc.log.Debugf("getting summoner by id %s", id)
	body, err := sc.get(fmt.Sprintf("/summoners/%s", id))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	res := &Summoner{}
	errDec := json.NewDecoder(body).Decode(res)
	if errDec != nil {
		sc.log.Errorf("error decoding summoner by id %s : %s", id, errDec.Error())
		return nil, &Error{ErrorDecode, errDec.Error()}
	}
	return res, nil
}

func (sc *summonerClient) get(url string) (io.ReadCloser, *Error) {
	return sc.c.Get(fmt.Sprintf("/summoner/v1%s", url))
}

type Summoner struct {
	AccountId     string `json:"accountId" bson:"accountId"`
	ProfileIconId int    `json:"profileIconId" bson:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate" bson:"revisionDate"`
	Name          string `json:"name" bson:"name"`
	Id            string `json:"id" bson:"id"`
	Puuid         string `json:"puuid" bson:"puuid"`
	SummonerLevel int    `json:"summonerLevel" bson:"summonerLevel"`
}
