package api

import (
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
)

var apiKey string
var log logger

func init() {
	file, _ := ioutil.ReadFile("../apikey")
	apiKey = string(file)
	testLogger := logrus.New()
	testLogger.SetLevel(logrus.DebugLevel)
	log = testLogger
}

func TestNewAPI(t *testing.T) {
	_ = NewAPI(apiKey, EUW)
}
