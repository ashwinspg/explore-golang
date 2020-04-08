package utils

import (
	"time"

	"github.com/sirupsen/logrus"
)

const LogRefKey = "reference"

//LogEntryWithRef returns a logrus Entry with a random unique value for requestId field
func LogEntryWithRef() *logrus.Entry {
	return logrus.WithField(LogRefKey, time.Now().Unix())
}
