package summoner

import (
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
