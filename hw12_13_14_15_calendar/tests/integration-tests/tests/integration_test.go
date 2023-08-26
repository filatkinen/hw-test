//go:build integration

package tests

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq" // import pq
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/require"
)

var (
	conf    ConfigEnv
	db      *sql.DB
	conn    *amqp.Connection
	channel *amqp.Channel
	chMsg   <-chan amqp.Delivery
	notices map[string]bool
)

func TestMain(m *testing.M) {
	//os.Setenv("DB_USER", "calendar")
	//os.Setenv("DB_PASS", "pass")
	//os.Setenv("DB_HOST", "localhost")
	//os.Setenv("DB_PORT", "5432")
	//os.Setenv("DB_NAME", "calendar")
	//
	//os.Setenv("S_HOST", "localhost")
	//os.Setenv("S_PORT", "8800")
	//
	//os.Setenv("R_USER", "rmuser")
	//os.Setenv("R_PASS", "rmpass")
	//os.Setenv("R_HOST", "localhost")
	//os.Setenv("R_PORT", "5672")
	//os.Setenv("E_EXCHANGE", "calendar-exchange")

	failOnError(setupTest(), "init testing")

	result := m.Run()
	failOnError(finishTest(), "finishing testing")
	os.Exit(result)
}

func failOnError(err error, msg string) {
	if err != nil {
		e := finishTest()
		err = errors.Join(e, err)
		log.Fatalf("%s: %s", msg, err)
	}
}

func setupTest() (err error) {
	err = envconfig.Process("notify_service", &conf)
	failOnError(err, "error reading env")

	DBString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&Timezone=UTC",
		conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)
	RabbitString := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		conf.RabbitUser, conf.RabbitPass, conf.RabbitHost, conf.RabbitPort)

	db, err = sql.Open("postgres", DBString)
	failOnError(err, "error while opening DB")
	err = db.PingContext(context.Background())
	failOnError(err, "error while ping DB")

	// rabbit
	conn, err = amqp.Dial(RabbitString)
	failOnError(err, "Could not connect to rabbit")

	channel, err = conn.Channel()
	failOnError(err, "Coud not reate channel to rabbit")

	err = channel.ExchangeDeclare(
		conf.RabbiExchange, // name
		"fanout",           // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Coud not declare Exchange rabbit")

	queue, err := channel.QueueDeclare(
		"",    // name of the queue
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Could not declare Queue rabbit")

	err = channel.QueueBind(
		queue.Name,         // queue name
		"",                 // routing key
		conf.RabbiExchange, // exchange
		false,
		nil,
	)
	failOnError(err, "Could not bind Queue rabbit")

	chMsg, err = channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	notices = make(map[string]bool)
	go func() {
		var notice storage.Notice
		for msg := range chMsg {
			err = json.Unmarshal(msg.Body, &notice)
			failOnError(err, "error unmarshaling notice")
			notices[notice.ID] = true
		}
	}()
	failOnError(err, "Could not create consume channel rabbit")
	return err
}

func finishTest() (err error) {
	if db != nil {
		_, err = db.Exec("truncate table events")
		if e := db.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}

	if channel != nil {
		if e := channel.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}
	if conn != nil {
		if e := conn.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}

func TestHttp(t *testing.T) {
	ServiceURL := fmt.Sprintf("http://%s:%s/", conf.ServiceHost, conf.ServicePort)

	// test status
	t.Run("test HTTP status", func(t *testing.T) {
		testHTTPStatus(t, ServiceURL)
	})

	// test API
	t.Run("test HTTP API + bisnes logic", func(t *testing.T) {
		testHTTPAPI(t, ServiceURL)
	})
}

func TestScheduleDeleteAndSendMessage(t *testing.T) {
	ServiceURL := fmt.Sprintf("http://%s:%s/", conf.ServiceHost, conf.ServicePort)

	UserID1, err := storage.UUID()
	if err != nil {
		log.Fatalf("error generating user UUID %s\n", err)
	}
	var events []*storage.Event
	basedate := time.Now().UTC().Add(-time.Hour*24*365 - time.Hour)
	// generating data
	for i := 0; i < 5; i++ {
		var event storage.Event
		event.DateTimeStart = basedate.Add(-time.Hour * 24 * time.Duration(i))
		event.DateTimeEnd = basedate.Add(-time.Hour*24*time.Duration(i) + time.Minute*30)
		event.Title = fmt.Sprintf("Task %d", i)
		event.UserID = UserID1
		events = append(events, &event)
	}

	client := http.Client{
		Timeout: time.Second * 2,
	}
	URL := fmt.Sprintf("%sevent", ServiceURL)

	// add events to delete
	for i := range events {
		fullURL := URL + "/add/" + events[i].UserID
		eventid, err := testAddEvent(t, client, events[i], fullURL)
		require.NoError(t, err)
		events[i].ID = eventid
	}

	// add one event to notice
	var event storage.Event
	basedate = time.Now().UTC().Add(time.Second + 2)
	event.DateTimeStart = basedate.Add(time.Minute * 30)
	event.DateTimeEnd = basedate.Add(time.Minute * 60)
	event.DateTimeNotice = basedate
	event.UserID = UserID1
	fullURL := URL + "/add/" + event.UserID
	EventIDRabbit, err := testAddEvent(t, client, &event, fullURL)
	require.NoError(t, err)

	// wait for the schedule delete old events and send notice to rabbit
	time.Sleep(time.Second * 5)

	// checking that events are deleted
	t.Run("Checking that old events(one year) deleted", func(t *testing.T) {
		for i := range events {
			fullURL := URL + "/get/" + UserID1
			eventID := events[i].ID
			testGetNoEvent(t, client, eventID, fullURL)
		}
	})

	t.Run("Checking that notice was delivered using rabbit", func(t *testing.T) {
		_, ok := notices[EventIDRabbit]
		require.True(t, ok, "Not Found Event ID in rabbit queue")
	})
}

func testHTTPStatus(t *testing.T, ServiceURL string) { //nolint
	URL := fmt.Sprintf("%sunknown", ServiceURL)
	resp, err := http.Get(URL) //nolint
	defer resp.Body.Close()    //nolint
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)

	resp, err = http.Get(ServiceURL) //nolint
	defer resp.Body.Close()          //nolint
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func testHTTPAPI(t *testing.T, ServiceURL string) { //nolint
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
	URL := fmt.Sprintf("%sevent", ServiceURL)

	// add events
	t.Run("Add Events", func(t *testing.T) {
		for i := range events {
			fullURL := URL + "/add/" + events[i].UserID
			eventid, err := testAddEvent(t, client, events[i], fullURL)
			require.NoError(t, err)
			events[i].ID = eventid
		}
	})

	t.Run("Change events + get events", func(t *testing.T) {
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
	})

	URL = fmt.Sprintf("%sevents", ServiceURL)
	userid := UserID2
	date := events[len(events)-1].DateTimeStart

	t.Run("List day", func(t *testing.T) {
		fullURL := URL + "/day/" + userid
		ev := testList(t, client, date, fullURL)
		require.Equal(t, 1, len(ev))
	})
	t.Run("List week", func(t *testing.T) {
		fullURL := URL + "/week/" + userid
		ev := testList(t, client, date, fullURL)
		require.Equal(t, 1, len(ev))
	})
	t.Run("List month", func(t *testing.T) {
		fullURL := URL + "/month/" + userid
		ev := testList(t, client, date, fullURL)
		require.Equal(t, 5, len(ev))
	})

	t.Run("Delete event", func(t *testing.T) {
		// delete event
		URL = fmt.Sprintf("%sevent", ServiceURL)
		userid = UserID2
		eventID := events[len(events)-1].ID
		fullURL := URL + "/delete/" + userid
		testDelete(t, client, eventID, fullURL)
	})

	t.Run("Try to get previously deleted event", func(t *testing.T) {
		fullURL := URL + "/get/" + userid
		eventID := events[len(events)-1].ID
		testGetNoEvent(t, client, eventID, fullURL)
	})
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
