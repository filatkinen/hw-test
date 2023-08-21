package sender

import (
	"os"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/rabbit"
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel       logger.LoggingLevel
	Logfile        string
	Rabbit         rabbit.Config
	ServiceLogfile string
}

func NewConfig(in string) (Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(in)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	config := Config{
		Logfile:        viper.GetString("logger.logfile"),
		LogLevel:       logger.GetLoggingLevel(viper.GetString("logger.loglevel")),
		ServiceLogfile: viper.GetString("service.logfile"),
		Rabbit: rabbit.Config{
			Port:     viper.GetString("rabbit.port"),
			Address:  viper.GetString("rabbit.address"),
			User:     os.Getenv(viper.GetString("env.rabbituser")),
			Password: os.Getenv(viper.GetString("env.rabbitpass")),
			Queue:    viper.GetString("rabbit.queue"),
			Tag:      viper.GetString("rabbit.tag"),
		},
	}
	return config, nil
}
