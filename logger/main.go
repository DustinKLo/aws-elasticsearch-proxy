package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logging = &logrus.Logger{
	Level: logrus.DebugLevel,
	Formatter: &prefixed.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		ForceFormatting: true,
	},
}

func init() {
	var logFile = os.Getenv("LOG_FILE_PATH")
	if logFile == "" {
		Logging.SetOutput(os.Stdout)
	} else {
		Logging.SetOutput(&lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    500, // MBs
			MaxBackups: 1,
			MaxAge:     28, // days
		})
	}

	logLevel := os.Getenv("LOG_LEVEL")
	switch strings.ToLower(logLevel) {
	case "error":
		Logging.SetLevel(logrus.ErrorLevel)
	case "warn":
		Logging.SetLevel(logrus.WarnLevel)
	case "warning":
		Logging.SetLevel(logrus.WarnLevel)
	case "debug":
		Logging.SetLevel(logrus.DebugLevel)
	default:
		Logging.SetLevel(logrus.WarnLevel)
	}
}
