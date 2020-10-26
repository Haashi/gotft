package gotft

type GoTFT struct {
	BR  *api
	EUN *api
	EUW *api
	JP  *api
	KR  *api
	LAN *api
	LAS *api
	NA  *api
	OCE *api
	TR  *api
	RU  *api
}

func New(apikey string, options ...Option) *GoTFT {
	gotft := &GoTFT{}
	gotft.BR = newAPI(apikey, BR, options...)
	gotft.EUN = newAPI(apikey, EUN, options...)
	gotft.EUW = newAPI(apikey, EUW, options...)
	gotft.JP = newAPI(apikey, JP, options...)
	gotft.KR = newAPI(apikey, KR, options...)
	gotft.LAN = newAPI(apikey, LAN, options...)
	gotft.LAS = newAPI(apikey, LAS, options...)
	gotft.NA = newAPI(apikey, NA, options...)
	gotft.OCE = newAPI(apikey, OCE, options...)
	gotft.RU = newAPI(apikey, RU, options...)
	gotft.TR = newAPI(apikey, TR, options...)
	return gotft
}
