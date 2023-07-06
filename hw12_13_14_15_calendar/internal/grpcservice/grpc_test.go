package grpcservice_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/grpcservice/client"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/server/http"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestGRPC(t *testing.T) {
	config := server.Config{
		LogLevel:        logger.LevelInfo,
		Logfile:         "",
		StoreType:       "memory",
		ServicePort:     "9091",
		ServiceGrpcPort: "50052",
		ServiceAddress:  "",
		ServiceLogfile:  "",
		DB:              server.DBConfig{},
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

	time.Sleep(time.Second)

	// test API
	t.Run("test GRPC API", func(t *testing.T) {
		testGRPCAPI(t, net.JoinHostPort(config.ServiceAddress, config.ServiceGrpcPort))
	})

	err = srv.Stop(context.Background())
	if err != nil {
		log.Printf("error stopping server %s", err)
	}
	<-chanEnd
}

func testGRPCAPI(t *testing.T, address string) { //nolint
	cl := client.NewGrpcClientCalendar(address)
	err := cl.Start()
	require.NoError(t, err)
	defer cl.Close()

	UserID1, err := storage.UUID()
	require.NoError(t, err, "error generating user UUID")
	UserID2, err := storage.UUID()
	require.NoError(t, err, "error generating user UUID")

	var events []*storage.Event
	basedate := time.Date(2023, 7, 4, 1, 0, 0, 0, time.UTC)

	// generating data
	for i := 0; i < 10; i++ {
		var event storage.Event
		event.DateTimeStart = basedate.Add(-time.Hour * 24 * time.Duration(i))
		event.DateTimeEnd = basedate.Add(-time.Hour*24*time.Duration(i) + time.Minute*30)
		event.Title = fmt.Sprintf("Task %d", i)
		if i < 5 {
			event.UserID = UserID1
		} else {
			event.UserID = UserID2
		}
		events = append(events, &event)
	}

	ctx := context.Background()
	// add events
	for i := range events {
		eventid, err := cl.AddEvent(ctx, events[i])
		require.NoError(t, err)
		events[i].ID = eventid
	}

	// change events + get events
	for i := range events {
		eventPut := *events[i]
		eventPut.Title += " Changed"
		err = cl.ChangeEventUser(ctx, &eventPut, eventPut.UserID)
		require.NoError(t, err)

		eventGet, err := cl.GetEvent(ctx, eventPut.ID)
		require.NoError(t, err)
		require.Equal(t, eventPut, *eventGet)
	}

	// list day, week, month
	userid := UserID2
	date := events[len(events)-1].DateTimeStart

	ev, err := cl.ListEventsDay(ctx, date, userid)
	require.NoError(t, err)
	require.Equal(t, 1, len(ev))

	ev, err = cl.ListEventsWeek(ctx, date, userid)
	require.NoError(t, err)
	require.Equal(t, 1, len(ev))

	ev, err = cl.ListEventsMonth(ctx, date, userid)
	require.NoError(t, err)
	require.Equal(t, 5, len(ev))

	// delete event
	userid = UserID2
	eventID := events[len(events)-1].ID
	err = cl.DeleteEventUser(ctx, eventID, userid)
	require.NoError(t, err)

	// get deleted event - got error
	err = cl.DeleteEventUser(ctx, eventID, userid)
	require.NotNil(t, err)
	log.Println(err)
}
