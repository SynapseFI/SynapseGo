package wrapper

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
