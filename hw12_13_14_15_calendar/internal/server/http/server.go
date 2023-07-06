package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	srv "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/grpcservice/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/server/calendar"
)

type Server struct {
	srv        *http.Server
	logSrv     *logger.Logger
	logHTTP    *httplog
	app        *calendar.App
	grpcserver *srv.GrpcServerCalendar
}

func NewServer(config server.Config, logSrv *logger.Logger) (*Server, error) {
	serv := Server{
		srv: &http.Server{
			Addr:              net.JoinHostPort(config.ServiceAddress, config.ServicePort),
			ReadHeaderTimeout: time.Second * 5,
		},
		logSrv: logSrv,
	}

	serv.srv.Handler = serv.rootHandler()

	app, err := calendar.New(logSrv, config)
	if err != nil {
		return nil, err
	}
	serv.app = app
	serv.logHTTP = newHTTPLogger(config.ServiceLogfile, logSrv)

	serv.grpcserver = srv.NewGrpcServerCalendar(serv.app, config, serv.logHTTP.httplogger)

	return &serv, nil
}

func (s *Server) Start() error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var err1, err2 error
	go func() {
		defer wg.Done()
		err1 = s.startHTTP()
		if err1 != nil {
			s.logSrv.Error("failed to start HTTP server: " + err1.Error())
		}
	}()
	go func() {
		defer wg.Done()
		err2 = s.startGRPC()
		if err2 != nil {
			s.logSrv.Error("failed to start GRPC server: " + err2.Error())
		}
	}()
	wg.Wait()
	if err1 != nil || err2 != nil {
		return errors.Join(err1, err2)
	}
	return nil
}

func (s *Server) startHTTP() error {
	s.logSrv.Info(fmt.Sprintf("Starting HTTP server: %s", s.srv.Addr))
	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) startGRPC() error {
	s.logSrv.Info(fmt.Sprintf("Starting GRPC server: %s", s.grpcserver.Addr))
	err := s.grpcserver.Start()
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	s.grpcserver.Close()
	s.logSrv.Info("GRPC  shutdown complete.")

	if err := s.srv.Shutdown(ctx); err != nil {
		s.logSrv.Error(fmt.Sprintf("HTTP shutdown error: %v", err))
		return err
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
