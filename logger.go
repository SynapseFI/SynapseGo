package synapse

import (
	logging "log"
)

type (
	logger struct {
	}
)

var log *logger

// Error used to log errors
func (l *logger) error(args ...interface{}) {
	if logMode {
		logging.Println(args...)
	}
}

// Info used to log information
func (l *logger) info(args ...interface{}) {
	if logMode {
		logging.Println(args...)
	}
}
