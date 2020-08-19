package gotft

import "github.com/Haashi/gotft/api"

type GoTFT struct {
	BR  *api.API
	EUW *api.API
	EUN *api.API
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
	BR  = "br"
	EUW = "euw"
	EUN = "eun"
	JP  = "jp"
	KR  = "kr"
	LAN = "lan"
	LAS = "las"
	NA  = "na"
	OCE = "oce"
	TR  = "tr"
	RU  = "ru"
)

func NewGOTFT(apikey string, prodKey bool) *GoTFT {
	gotft := &GoTFT{}
	gotft.BR = api.NewAPI(apikey, BR, prodKey)
	gotft.EUW = api.NewAPI(apikey, EUW, prodKey)
	gotft.EUN = api.NewAPI(apikey, EUN, prodKey)
	gotft.JP = api.NewAPI(apikey, JP, prodKey)
	gotft.KR = api.NewAPI(apikey, KR, prodKey)
	gotft.LAN = api.NewAPI(apikey, LAN, prodKey)
	gotft.LAS = api.NewAPI(apikey, LAS, prodKey)
	gotft.NA = api.NewAPI(apikey, NA, prodKey)
	gotft.OCE = api.NewAPI(apikey, OCE, prodKey)
	gotft.TR = api.NewAPI(apikey, TR, prodKey)
	gotft.RU = api.NewAPI(apikey, RU, prodKey)
	return gotft
}
