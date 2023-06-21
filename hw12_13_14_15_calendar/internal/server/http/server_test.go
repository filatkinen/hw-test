package internalhttp_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/server/http"
	"github.com/stretchr/testify/require"
)

func TestHttpStatus(t *testing.T) {
	config := server.Config{
		LogLevel:       logger.LevelInfo,
		Logfile:        "",
		StoreType:      "memory",
		ServicePort:    "9090",
		ServiceAddress: "",
		ServiceLogfile: "",
		DB:             server.DBConfig{},
	}
	l := logger.New(config.LogLevel, os.Stdout)

	srv, err := internalhttp.NewServer(config, l)
	if err != nil {
		log.Fatalf("error crating http server %s", err)
	}
	defer srv.Close()

	chanEnd := make(chan struct{})
	go func() {
		errstop := srv.Start()
		if errstop != nil {
			log.Fatalf("error starting server %s", errstop)
		}
		chanEnd <- struct{}{}
	}()
	URL := fmt.Sprintf("http://localhost:%s/unknown", config.ServicePort)
	resp, err := http.Get(URL) //nolint
	defer resp.Body.Close()    //nolint
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)

	URL = fmt.Sprintf("http://localhost:%s/", config.ServicePort)
	resp, err = http.Get(URL) //nolint
	defer resp.Body.Close()   //nolint
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	err = srv.Stop(context.Background())
	if err != nil {
		log.Printf("error stopping server %s", err)
	}
	<-chanEnd
}
