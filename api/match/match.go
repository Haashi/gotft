package match

import (
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
