package gotft

import (
	"testing"

	"github.com/haashi/gotft/internal"
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
	_ = New(apiKey, WithClient(internal.NewDefaultClient()), WithLog(logrus.New()))
}
