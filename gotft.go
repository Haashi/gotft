package gotft

import "github.com/Haashi/gotft/api"

type GoTFT struct {
	BR  *api.API
	EUN *api.API
	EUW *api.API
	JP  *api.API
	KR  *api.API
	LAN *api.API
	LAS *api.API
	NA  *api.API
	OCE *api.API
	TR  *api.API
	RU  *api.API
}

const (
	BR  = "br1"
	EUN = "eun1"
	EUW = "euw1"
	JP  = "jp1"
	KR  = "kr"
	LAN = "la1"
	LAS = "la2"
	NA  = "na1"
	OCE = "oc1"
	RU  = "ru"
	TR  = "tr1"
)

func NewGOTFT(apikey string) *GoTFT {
	gotft := &GoTFT{}
	gotft.BR = api.NewAPI(apikey, BR)
	gotft.EUN = api.NewAPI(apikey, EUN)
	gotft.EUW = api.NewAPI(apikey, EUW)
	gotft.JP = api.NewAPI(apikey, JP)
	gotft.KR = api.NewAPI(apikey, KR)
	gotft.LAN = api.NewAPI(apikey, LAN)
	gotft.LAS = api.NewAPI(apikey, LAS)
	gotft.NA = api.NewAPI(apikey, NA)
	gotft.OCE = api.NewAPI(apikey, OCE)
	gotft.RU = api.NewAPI(apikey, RU)
	gotft.TR = api.NewAPI(apikey, TR)
	return gotft
}
