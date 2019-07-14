package gormrus

import "github.com/sirupsen/logrus"

type Logger struct{}

func (l *Logger) Print(values ...interface{}) {
	if len(values) <= 1 {
		return
	}

	logrus.WithFields(logrus.Fields{
		"module": "gorm",
		"type":   values[0],
	}).Info(values[1:])
}
