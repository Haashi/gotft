package gotft

import "github.com/haashi/gotft/api"

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

func NewGOTFT(apikey string, options ...api.Option) *GoTFT {
	gotft := &GoTFT{}
	gotft.BR = api.NewAPI(apikey, api.BR, options...)
	gotft.EUN = api.NewAPI(apikey, api.EUN, options...)
	gotft.EUW = api.NewAPI(apikey, api.EUW, options...)
	gotft.JP = api.NewAPI(apikey, api.JP, options...)
	gotft.KR = api.NewAPI(apikey, api.KR, options...)
	gotft.LAN = api.NewAPI(apikey, api.LAN, options...)
	gotft.LAS = api.NewAPI(apikey, api.LAS, options...)
	gotft.NA = api.NewAPI(apikey, api.NA, options...)
	gotft.OCE = api.NewAPI(apikey, api.OCE, options...)
	gotft.RU = api.NewAPI(apikey, api.RU, options...)
	gotft.TR = api.NewAPI(apikey, api.TR, options...)
	return gotft
}
