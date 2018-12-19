package synapse

import (
	logrus "github.com/sirupsen/logrus"
)

type (
	logger struct {
	}
)

var log *logger

// Info used to log information
func (l *logger) info(args ...interface{}) {
	if developerMode {
		logrus.Info(args...)
	}
}

// Error used to log errors
func (l *logger) error(args ...interface{}) {
	if developerMode {
		logrus.Error(args...)
	}
}
