package utils

import (
	"github.com/sirupsen/logrus"	
)

// LogErrorf logs a message at level Error on the standard logger.
func LogErrorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// LogError logs a message at level Error on the standard logger.
func LogError(args ...interface{}) {
	logrus.Error(args...)
}
