package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/scheduler"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/rabbit/producer"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/scheduler_config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	conf, err := scheduler.NewConfig(configFile)
	if err != nil {
		log.Fatalf("error reading config file %v", err)
	}

	flog := os.Stdout
	if len(conf.Logfile) != 0 {
		f, err := os.OpenFile(conf.Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
		if err == nil {
			flog = f
		}
	}
	l := logger.New(conf.LogLevel, flog)
	if flog == os.Stdout && len(conf.Logfile) != 0 {
		l.Info(fmt.Sprintf("Error opening file %s for logging. Using console", conf.Logfile))
	}
	defer l.Close()

	schedul, err := producer.NewProducer(conf, l)
	if err != nil {
		l.Error("failed to make new app server due error: " + err.Error())
		return
	}
	defer func() {
		if err := schedul.Close(); err != nil {
			l.Error(err.Error())
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		schedul.Start()
	}()

	<-signalCh
	schedul.Stop()
}
