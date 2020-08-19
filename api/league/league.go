package league

import (
	"github.com/Haashi/gotft/api/client"
)

type LeagueEndpoint struct {
	c *client.Client
}

func (l *LeagueEndpoint) GetMasterLeague() {
	l.c.Get("/league/v1/master")
}

func NewLeagueEndpoint(c *client.Client) *LeagueEndpoint {
	le := &LeagueEndpoint{}
	le.c = c
	return le
}
