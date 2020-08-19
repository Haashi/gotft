package match

import (
	"testing"

	"github.com/Haashi/gotft/api/client"
)

func TestNewMatchEndpoint(t *testing.T) {
	c := client.NewClient("RGAPI-2b5f8fc6-5ce2-4736-8bf2-731b8b29c376", "euw", false)
	_ = NewMatchEndpoint(c)
}

func TestMatchEndpoint_GetByPUUID(t *testing.T) {
	c := client.NewClient("RGAPI-2b5f8fc6-5ce2-4736-8bf2-731b8b29c376", "euw", false)
	me := NewMatchEndpoint(c)
	matches := me.GetMatchesByPUUID("5_QD37vWUa7Eq8jP6Cy-R18z60E9nRJlpkDZjqBGvtngedjANG6221udHyYnN2wCJCZV7CnlAqcnHQ")
	if len(*matches) == 0 {
		t.FailNow()
	}
}
