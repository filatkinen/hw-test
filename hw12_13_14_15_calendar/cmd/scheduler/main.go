package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/scheduler"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/rabbit/producer"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/server/calendar"
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

	app, err := calendar.New(l, server.Config{
		LogLevel:  conf.LogLevel,
		StoreType: conf.StoreType,
		DB:        conf.DB,
	})
	if err != nil {
		l.Error("failed to make new app server due error: " + err.Error())
		return
	}
	defer app.Close(context.Background())

	schedul, err := producer.NewProducer(conf, l)
	if err != nil {
		l.Error("failed to make new app producer due error: " + err.Error())
		return
	}
	defer func() {
		if err := schedul.Close(); err != nil {
			l.Error(err.Error())
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	dofunc := func() [][]byte {
		DeleteOldEvents(app, l)
		return GetEventsToSend(app, l)
	}

	go func() {
		schedul.Start(dofunc)
	}()

	<-signalCh
	schedul.Stop()
}

func DeleteOldEvents(app *calendar.App, log *logger.Logger) {
	ctx := context.Background()
	count, err := app.DeleteOldEvents(ctx, time.Now().UTC())
	if err != nil {
		log.Error("Error while getting events to delete: " + err.Error())
		return
	}
	if count != 0 {
		log.Logging(logger.LevelInfo, fmt.Sprintf("Deleted %d old events", count))
	}
}

func GetEventsToSend(app *calendar.App, log *logger.Logger) [][]byte {
	ctx := context.Background()
	notices, err := app.ListNoticesToSend(ctx, time.Now().UTC())
	if err != nil {
		log.Error("Error while getting events to notice: " + err.Error())
		return nil
	}
	messages := make([][]byte, 0, len(notices))
	for i := range notices {
		b, err := json.Marshal(*notices[i])
		if err != nil {
			log.Error("Error while serializing: " + err.Error())
			return nil
		}
		messages = append(messages, b)
	}
	return messages
}
