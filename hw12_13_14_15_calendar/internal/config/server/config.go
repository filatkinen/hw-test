package server

import (
	"os"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/spf13/viper"
)

const DefaultMaxIdleTime time.Duration = time.Minute * 15

type DBConfig struct {
	DBUser       string
	DBPass       string
	DBAddress    string
	DBPort       string
	DBName       string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}

type Config struct {
	LogLevel       logger.LoggingLevel
	Logfile        string
	StoreType      string
	ServicePort    string
	ServiceAddress string
	ServiceLogfile string
	DB             DBConfig
}

func NewConfig(in string) (Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(in)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	viper.SetDefault("db.maxopenconns", 10)
	viper.SetDefault("db.maxidleconns", 10)
	viper.SetDefault("db.maxidletime", DefaultMaxIdleTime.String())

	duration, err := time.ParseDuration(viper.GetString("db.maxidletime"))
	if err != nil {
		duration = DefaultMaxIdleTime
	}

	config := Config{
		Logfile:        viper.GetString("logger.logfile"),
		LogLevel:       logger.GetLoggingLevel(viper.GetString("logger.loglevel")),
		StoreType:      viper.GetString("storage.type"),
		ServicePort:    viper.GetString("bindings.port"),
		ServiceAddress: viper.GetString("bindings.address"),
		ServiceLogfile: viper.GetString("service.logfile"),
		DB: DBConfig{
			DBUser:       os.Getenv(viper.GetString("env.dbuser")),
			DBPass:       os.Getenv(viper.GetString("env.dbpass")),
			DBAddress:    viper.GetString("db.address"),
			DBPort:       viper.GetString("db.port"),
			DBName:       viper.GetString("db.db"),
			MaxOpenConns: viper.GetInt("db.maxopenconns"),
			MaxIdleConns: viper.GetInt("db.maxidleconns"),
			MaxIdleTime:  duration,
		},
	}

	return config, nil
}
