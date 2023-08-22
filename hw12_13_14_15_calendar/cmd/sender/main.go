package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/sender"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/rabbit/consumer"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/sender_config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	conf, err := sender.NewConfig(configFile)
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

	senderer, err := consumer.NewConsumer(conf, l)
	if err != nil {
		l.Error("failed to make new app server due error: " + err.Error())
		return
	}
	defer func() {
		if err := senderer.Close(); err != nil {
			l.Error(err.Error())
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	dofunc := func(message []byte) {
		var notice storage.Notice
		err := json.Unmarshal(message, &notice)
		if err != nil {
			l.Error(fmt.Sprintf("Error while deserialosing message: %s", err))
			return
		}
		l.Logging(logger.LevelInfo, fmt.Sprintf("Got message from rabbit and log it. EventID:%s DateTimeStart:%s",
			notice.ID, notice.DateTime))
	}

	go func() {
		senderer.Start(dofunc)
	}()

	<-signalCh
	senderer.Stop()
}
