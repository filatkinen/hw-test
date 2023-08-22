package internalhttp_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/server/http"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestHttp(t *testing.T) {
	config := server.Config{
		LogLevel:        logger.LevelInfo,
		Logfile:         "",
		StoreType:       "memory",
		ServicePort:     "9090",
		ServiceGrpcPort: "50051",
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
	// test status
	time.Sleep(time.Second)
	t.Run("test HTTP status", func(t *testing.T) {
		testHTTPStatus(t, config.ServicePort)
	})

	// test API
	t.Run("test HTTP API", func(t *testing.T) {
		testHTTPAPI(t, config.ServicePort)
	})

	err = srv.Stop(context.Background())
	if err != nil {
		log.Printf("error stopping server %s", err)
	}
	<-chanEnd
}

func testHTTPStatus(t *testing.T, servicePort string) { //nolint
	URL := fmt.Sprintf("http://localhost:%s/unknown", servicePort)
	resp, err := http.Get(URL) //nolint
	defer resp.Body.Close()    //nolint
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)

	URL = fmt.Sprintf("http://localhost:%s/", servicePort)
	resp, err = http.Get(URL) //nolint
	defer resp.Body.Close()   //nolint
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func testHTTPAPI(t *testing.T, servicePort string) { //nolint
	UserID1, err := storage.UUID()
	if err != nil {
		log.Fatalf("error generating user UUID %s\n", err)
	}
	UserID2, err := storage.UUID()
	if err != nil {
		log.Fatalf("error generating user UUID %s\n", err)
	}
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

	client := http.Client{
		Timeout: time.Second * 2,
	}
	URL := fmt.Sprintf("http://localhost:%s/event", servicePort)

	// add events
	for i := range events {
		fullURL := URL + "/add/" + events[i].UserID
		eventid, err := testAddEvent(t, client, events[i], fullURL)
		require.NoError(t, err)
		events[i].ID = eventid
	}

	// change events + get events
	for i := range events {
		fullURL := URL + "/change/" + events[i].UserID
		eventPut := *events[i]
		eventPut.Title += " Changed"
		testChangeEvent(t, client, &eventPut, fullURL)

		fullURL = URL + "/get/" + events[i].UserID
		eventGet := testGetEvent(t, client, eventPut.ID, fullURL)
		require.Equal(t, eventPut, *eventGet)
	}

	// list day, week, month
	URL = fmt.Sprintf("http://localhost:%s/events", servicePort)
	userid := UserID2
	date := events[len(events)-1].DateTimeStart
	fullURL := URL + "/day/" + userid
	ev := testList(t, client, date, fullURL)
	require.Equal(t, 1, len(ev))
	fullURL = URL + "/week/" + userid
	ev = testList(t, client, date, fullURL)
	require.Equal(t, 1, len(ev))
	fullURL = URL + "/month/" + userid
	ev = testList(t, client, date, fullURL)
	require.Equal(t, 5, len(ev))

	// delete event
	URL = fmt.Sprintf("http://localhost:%s/event", servicePort)
	userid = UserID2
	eventID := events[len(events)-1].ID
	fullURL = URL + "/delete/" + userid
	testDelete(t, client, eventID, fullURL)

	// get deleted event - 404 status
	fullURL = URL + "/get/" + userid
	testGetNoEvent(t, client, eventID, fullURL)
}

func testDelete(t *testing.T, client http.Client, eventID string, url string) { //nolint
	data, err := json.Marshal(eventID)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	require.NoError(t, err)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func testList(t *testing.T, client http.Client, date time.Time, url string) []*storage.Event { //nolint
	data, err := json.Marshal(date)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data)) //nolint
	require.NoError(t, err)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	var ev []*storage.Event
	jsDecoder := json.NewDecoder(resp.Body)
	err = jsDecoder.Decode(&ev)
	require.NoError(t, err)
	return ev
}

func testAddEvent(t *testing.T, client http.Client, event *storage.Event, url string) (string, error) { //nolint
	data, err := json.Marshal(event)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	require.NoError(t, err)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var uid string
	jsDecoder := json.NewDecoder(resp.Body)
	err = jsDecoder.Decode(&uid)
	return uid, err
}

func testChangeEvent(t *testing.T, client http.Client, event *storage.Event, url string) { //nolint
	data, err := json.Marshal(event)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	require.NoError(t, err)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func testGetEvent(t *testing.T, client http.Client, eventID string, url string) *storage.Event { //nolint
	data, err := json.Marshal(eventID)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	require.NoError(t, err)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	var event storage.Event
	jsDecoder := json.NewDecoder(resp.Body)
	err = jsDecoder.Decode(&event)
	require.NoError(t, err)
	return &event
}

func testGetNoEvent(t *testing.T, client http.Client, eventID string, url string) { //nolint
	data, err := json.Marshal(eventID)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data)) //nolint
	require.NoError(t, err)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
