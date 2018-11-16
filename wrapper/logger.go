package wrapper

import (
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// Info used to log information
func infoLog(args ...interface{}) {
	if developerMode {
		logger.Info(args...)
	}
}

// Error used to log errors
func errorLog(args ...interface{}) {
	if developerMode {
		logger.Error(args...)
	}
}
