package summoner

import (
	"testing"

	"github.com/Haashi/gotft/api/client"
)

func TestNewSummonerEndpoint(t *testing.T) {
	c := client.NewClient("RGAPI-2b5f8fc6-5ce2-4736-8bf2-731b8b29c376", "euw", false)
	_ = NewSummonerEndpoint(c)
}

func TestSummonerEndpoint_GetByName(t *testing.T) {
	c := client.NewClient("RGAPI-2b5f8fc6-5ce2-4736-8bf2-731b8b29c376", "euw", false)
	se := NewSummonerEndpoint(c)
	summoner := se.GetByName("Haashi")
	if summoner.AccountId == "" {
		t.FailNow()
	}
}
