package league

import (
	"testing"

	"github.com/Haashi/gotft/api/client"
)

func TestNewLeagueEndpoint(t *testing.T) {
	c := client.NewClient("RGAPI-2b5f8fc6-5ce2-4736-8bf2-731b8b29c376", "euw", false)
	_ = NewLeagueEndpoint(c)
}

func TestLeagueEndpoint_GetMasterLeague(t *testing.T) {
	c := client.NewClient("RGAPI-2b5f8fc6-5ce2-4736-8bf2-731b8b29c376", "euw", false)
	le := NewLeagueEndpoint(c)
	master := le.GetMasterLeague()
	if master.LeagueId == "" {
		t.FailNow()
	}
}

func TestLeagueEndpoint_GetGrandmasterLeague(t *testing.T) {
	c := client.NewClient("RGAPI-2b5f8fc6-5ce2-4736-8bf2-731b8b29c376", "euw", false)
	le := NewLeagueEndpoint(c)
	grandmaster := le.GetGrandmasterLeague()
	if grandmaster.LeagueId == "" {
		t.FailNow()
	}
}

func TestLeagueEndpoint_GetChallengerLeague(t *testing.T) {
	c := client.NewClient("RGAPI-2b5f8fc6-5ce2-4736-8bf2-731b8b29c376", "euw", false)
	le := NewLeagueEndpoint(c)
	challenger := le.GetChallengerLeague()
	if challenger.LeagueId == "" {
		t.FailNow()
	}
}
