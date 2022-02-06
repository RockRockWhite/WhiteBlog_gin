package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func InitLogger(logFile string, level logrus.Level, timestampFormat string) {
	// 配置log
	file, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	logger = logrus.New()
	logger.Out = file
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: timestampFormat})
}

func Logger() *logrus.Logger {
	return logger
}
