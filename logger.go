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
	if logMode {
		logrus.Info(args...)
	}
}

// Error used to log errors
func (l *logger) error(args ...interface{}) {
	if logMode {
		logrus.Error(args...)
	}
}
