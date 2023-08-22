package scheduler

import (
	"log"
	"os"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/rabbit"
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel  logger.LoggingLevel
	Logfile   string
	StoreType string
	Rabbit    rabbit.Config
	DB        server.DBConfig
}

func NewConfig(in string) (Config, error) {
	const DefaultMaxIdleTime time.Duration = time.Minute * 15
	viper.SetConfigType("yaml")
	viper.SetConfigFile(in)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	viper.SetDefault("db.maxopenconns", 10)
	viper.SetDefault("db.maxidleconns", 10)
	viper.SetDefault("db.maxidletime", DefaultMaxIdleTime.String())

	duration, err := time.ParseDuration(viper.GetString("rabbit.checkinterval"))
	if err != nil {
		log.Printf("Error parsing db.maxidletime value: %s, using defaul tvalue:%s", err.Error(), DefaultMaxIdleTime)
		duration = rabbit.DefaultChecKTimeSheduler
	}
	config := Config{
		Logfile:   viper.GetString("logger.logfile"),
		LogLevel:  logger.GetLoggingLevel(viper.GetString("logger.loglevel")),
		StoreType: viper.GetString("storage.type"),
		DB: server.DBConfig{
			DBUser:       os.Getenv(viper.GetString("env.dbuser")),
			DBPass:       os.Getenv(viper.GetString("env.dbpass")),
			DBAddress:    viper.GetString("db.address"),
			DBPort:       viper.GetString("db.port"),
			DBName:       viper.GetString("db.db"),
			MaxOpenConns: viper.GetInt("db.maxopenconns"),
			MaxIdleConns: viper.GetInt("db.maxidleconns"),
			MaxIdleTime:  duration,
		},
		Rabbit: rabbit.Config{
			Port:          viper.GetString("rabbit.port"),
			Address:       viper.GetString("rabbit.address"),
			User:          os.Getenv(viper.GetString("env.rabbituser")),
			Password:      os.Getenv(viper.GetString("env.rabbitpass")),
			Queue:         viper.GetString("rabbit.queue"),
			CheckInterval: duration,
		},
	}
	return config, nil
}
