package main

import (
	"github.com/arisro/talos/cmd"
	"github.com/spf13/viper"
	"github.com/Sirupsen/logrus"
)

func main() {
	viper.SetDefault("log_level", "error")
	viper.SetEnvPrefix("talos")
	viper.BindEnv("log_level")

	logLevel := viper.Get("log_level")
	switch logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
		break
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
		break
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
		break
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
		break
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
		break
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
		break
	}

	cmd.Execute();
}
