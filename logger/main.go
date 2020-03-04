package logger

import (
	"os"
	"strings"

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

	switch strings.ToLower(configs.LogLevel) {
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
