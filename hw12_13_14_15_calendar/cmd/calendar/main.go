package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/server.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

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

	serv, err := internalhttp.NewServer(conf, l)
	if err != nil {
		l.Error("failed to make new app server due error: " + err.Error())
		return
	}
	defer func() {
		if err := serv.Close(); err != nil {
			l.Error(err.Error())
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	signalFailedStart := make(chan struct{})
	go func() {
		if err := serv.Start(); err != nil {
			l.Error("failed to start http server: " + err.Error())
		}
		signalFailedStart <- struct{}{}
	}()

	select {
	case <-signalFailedStart:
	case <-signalCh:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		if err := serv.Stop(ctx); err != nil {
			l.Error("failed to stop http server: " + err.Error())
		}
	}
}
