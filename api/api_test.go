package api

import (
	"io/ioutil"
	"testing"
)

var apiKey string

func init() {
	file, _ := ioutil.ReadFile("../apikey")
	apiKey = string(file)
}

func TestNewAPI(t *testing.T) {
	_ = NewAPI(apiKey, EUW)
}
