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
	gotft.BR = api.New(apikey, api.BR, options...)
	gotft.EUN = api.New(apikey, api.EUN, options...)
	gotft.EUW = api.New(apikey, api.EUW, options...)
	gotft.JP = api.New(apikey, api.JP, options...)
	gotft.KR = api.New(apikey, api.KR, options...)
	gotft.LAN = api.New(apikey, api.LAN, options...)
	gotft.LAS = api.New(apikey, api.LAS, options...)
	gotft.NA = api.New(apikey, api.NA, options...)
	gotft.OCE = api.New(apikey, api.OCE, options...)
	gotft.RU = api.New(apikey, api.RU, options...)
	gotft.TR = api.New(apikey, api.TR, options...)
	return gotft
}
