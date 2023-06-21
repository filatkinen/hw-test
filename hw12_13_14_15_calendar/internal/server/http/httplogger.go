package internalhttp

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
)

type httplog struct {
	httplogger *log.Logger
	fileLog    io.WriteCloser
}

func newHTTPLogger(fileNameLogHTTP string, logServer *logger.Logger) *httplog {
	httplogger := log.New(os.Stdout, "", 0)
	fileLog := os.Stdout
	f, err := os.OpenFile(fileNameLogHTTP, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err == nil {
		fileLog = f
		httplogger.SetOutput(fileLog)
	}
	if fileLog == os.Stdout {
		logServer.Error(fmt.Sprintf("Error openening file %s for HTTP logging, using os.Stdout", fileNameLogHTTP))
	} else {
		logServer.Info(fmt.Sprintf("logging HTTTP using %s", fileNameLogHTTP))
	}
	return &httplog{
		httplogger: httplogger,
		fileLog:    fileLog,
	}
}

func (s *httplog) close() error {
	if s.fileLog != os.Stdout {
		return s.fileLog.Close()
	}
	return nil
}
