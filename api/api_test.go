package api

import (
	"testing"

	"github.com/haashi/gotft/api/internal"
	"github.com/sirupsen/logrus"
)

var apiKey string
var testOpt *Options

func init() {
	apiKey = "apikey"
	testLogger := logrus.New()
	testLogger.SetLevel(logrus.DebugLevel)
	testOpt = &Options{log: testLogger, c: internal.NewDefaultClient()}
}

func TestNewAPI(t *testing.T) {
	_ = New(apiKey, EUW, WithClient(internal.NewDefaultClient()), WithLog(logrus.New()))
}
