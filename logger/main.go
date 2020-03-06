package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/hysds/aws-elasticsearch-proxy/configs"
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
	if configs.LogFile == "" {
		Logging.SetOutput(os.Stdout)
	} else {
		Logging.SetOutput(&lumberjack.Logger{
			Filename:   configs.LogFile,
			MaxSize:    500, // MBs
			MaxBackups: 1,
			MaxAge:     28, // days
		})
	}

	// ERROR 40
	// WARNING 30
	// INFO 20
	// DEBUG 10
	switch configs.LogLevel {
	case 40:
		Logging.SetLevel(logrus.ErrorLevel)
	case 30:
		Logging.SetLevel(logrus.WarnLevel)
	case 20:
		Logging.SetLevel(logrus.InfoLevel)
	case 10:
		Logging.SetLevel(logrus.DebugLevel)
	default:
		Logging.SetLevel(logrus.WarnLevel)
	}
}
