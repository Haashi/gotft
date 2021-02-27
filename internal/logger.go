package internal

import "github.com/sirupsen/logrus"

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

func NewDefaultLogger(fields logrus.Fields) Logger {
	defaultLogger := logrus.New()
	defaultLogger.SetLevel(logrus.InfoLevel)
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	defaultLogger.SetFormatter(formatter)
	entry := defaultLogger.WithFields(fields)
	return entry
}
