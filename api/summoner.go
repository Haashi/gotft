package api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/haashi/gotft/api/internal"
)

type summonerClient struct {
	c   *apiclient
	log internal.Logger
}

func newSummonerClient(c *apiclient, opt *Options) *summonerClient {
	opt.log.Debug("initializing summoner client")
	sc := &summonerClient{}
	sc.c = c
	sc.log = opt.log
	return sc
}

func (sc *summonerClient) GetByName(name string) (*Summoner, error) {
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
		return nil, ErrorDecode{fmt.Sprintf("summoner by name %s", name), errDec.Error()}
	}
	return res, nil
}

func (sc *summonerClient) GetByAccountId(accountId string) (*Summoner, error) {
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
		return nil, ErrorDecode{fmt.Sprintf("summoner by accountid %s", accountId), errDec.Error()}
	}
	return res, nil
}

func (sc *summonerClient) GetByPuuid(puuid string) (*Summoner, error) {
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
		return nil, ErrorDecode{fmt.Sprintf("summoner by puuid %s", puuid), errDec.Error()}
	}
	return res, nil
}

func (sc *summonerClient) GetById(id string) (*Summoner, error) {
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
		return nil, ErrorDecode{fmt.Sprintf("summoner  by id %s", id), errDec.Error()}
	}
	return res, nil
}

func (sc *summonerClient) get(url string) (io.ReadCloser, error) {
	return sc.c.get(fmt.Sprintf("/summoner/v1%s", url))
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
