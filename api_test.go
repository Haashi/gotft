package gotft

import (
	"io/ioutil"
	"testing"

	"github.com/haashi/gotft/internal"
	"github.com/sirupsen/logrus"
)

var apiKey string
var testOpt *Options

func init() {
	file, _ := ioutil.ReadFile("../apikey")
	apiKey = string(file)
	testLogger := logrus.New()
	testLogger.SetLevel(logrus.DebugLevel)
	testOpt = &Options{log: testLogger, c: internal.NewDefaultClient()}
}

func TestNewAPI(t *testing.T) {
	_ = newAPI(apiKey, EUW)
}
