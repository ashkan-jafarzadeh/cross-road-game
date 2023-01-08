package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

func GetLogger() *logrus.Logger {
	var log = logrus.New()
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		//Log.Formatter = &logrus.JSONFormatter{}
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr for now")
	}

	return log
}
