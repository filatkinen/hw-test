package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/server/calendar"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/calendar_config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	conf, err := server.NewConfig(configFile)
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

	app, err := calendar.New(l, conf)
	if err != nil {
		l.Error(fmt.Sprintf("Error creating  app %s", err.Error()))
		return
	}

	defer func() {
		if err := app.Close(context.Background()); err != nil {
			l.Error(err.Error())
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		ctx := context.Background()
		timer := time.NewTicker(time.Second * 2)
		defer timer.Stop()
		var counter int
		for {
			select {
			case <-signalCh:
				return
			case <-timer.C:
				t1 := time.Now().UTC().Truncate(time.Second)

				counter++
				eventID, _ := storage.UUID()
				userID, _ := storage.UUID()

				eventDelete := storage.Event{
					ID:             eventID,
					Title:          "Title Delete " + strconv.Itoa(counter),
					Description:    "Description Delete " + strconv.Itoa(counter),
					DateTimeStart:  t1.Add(-time.Hour * 24 * 365),
					DateTimeEnd:    t1.Add(-time.Hour*24*365 + time.Minute*30),
					DateTimeNotice: t1.Add(-time.Hour*24*365 + time.Minute*60),
					UserID:         userID,
				}
				_, err = app.AddEvent(ctx, &eventDelete, userID)
				if err != nil {
					l.Error(fmt.Sprintf("Error adding Delete  event %s", err.Error()))
				} else {
					l.Logging(logger.LevelInfo, fmt.Sprintf("Adding Events to DB (Delete): EventID:%s Time to start:%s",
						eventDelete.ID, eventDelete.DateTimeStart))
				}

				counter++
				eventID, _ = storage.UUID()
				userID, _ = storage.UUID()

				eventNotice := storage.Event{
					ID:             eventID,
					Title:          "Title Notice " + strconv.Itoa(counter),
					Description:    "Description Notice " + strconv.Itoa(counter),
					DateTimeStart:  t1.Add(time.Minute * 10),
					DateTimeEnd:    t1.Add(time.Minute * 30),
					DateTimeNotice: t1.Add(time.Second * 5),
					UserID:         userID,
				}
				_, err = app.AddEvent(ctx, &eventNotice, userID)
				if err != nil {
					l.Error(fmt.Sprintf("Error adding Notice  event %s", err.Error()))
				} else {
					l.Logging(logger.LevelInfo, fmt.Sprintf("Adding Events to DB (Notice): EventID:%s Time to start:%s",
						eventDelete.ID, eventDelete.DateTimeStart))
				}
			}
		}
	}()
	<-signalCh
}
