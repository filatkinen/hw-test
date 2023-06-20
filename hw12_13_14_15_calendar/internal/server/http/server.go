package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/server/calendar"
)

type Server struct {
	srv     *http.Server
	logSrv  *logger.Logger
	logHTTP *httplog
	app     *calendar.App
}

func NewServer(config server.Config, logSrv *logger.Logger) (*Server, error) {
	serv := Server{
		srv: &http.Server{
			Addr:              net.JoinHostPort(config.ServiceAddress, config.ServicePort),
			ReadHeaderTimeout: time.Second * 5,
		},
		logSrv: logSrv,
	}

	serv.srv.Handler = serv.routes()

	app, err := calendar.New(logSrv, config)
	if err != nil {
		return nil, err
	}
	serv.app = app

	serv.logHTTP = newHTTPLogger(config.ServiceLogfile, logSrv)

	return &serv, nil
}

func (s *Server) Start() error {
	s.logSrv.Info(fmt.Sprintf("Starting server: %s", s.srv.Addr))
	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		s.logSrv.Error(fmt.Sprintf("HTTP shutdown error: %v", err))
	}
	s.logSrv.Info("HTTP graceful shutdown complete.")

	return nil
}

func (s *Server) Close() error {
	err := s.logHTTP.close()
	if err != nil {
		s.logSrv.Error("Error closing HTTP log")
	}
	return s.app.Close(context.Background())
}
