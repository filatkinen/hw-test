package calendar

import (
	"context"
	"errors"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage/memory"
	mysqlstorage "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage/mysql"
	pgsqlstorage "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage/pqsql"
)

type App struct {
	servLog *logger.Logger
	storage storage.Store
}

func New(servLog *logger.Logger, config server.Config) (*App, error) {
	stor, err := newStorage(config)
	if err != nil {
		return nil, err
	}
	servLog.Info("application Calendar started")
	servLog.Info("application Calendar using db:" + config.StoreType)

	return &App{
		servLog: servLog,
		storage: stor,
	}, nil
}

func newStorage(config server.Config) (storage.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	switch config.StoreType {
	case "memory":
		return memorystorage.New(), nil
	case "mysql":
		stor, err := mysqlstorage.New(config)
		if err != nil {
			return nil, err
		}
		err = stor.Connect(ctx)
		if err != nil {
			return nil, err
		}
		return stor, err
	case "pgsql":
		stor, err := pgsqlstorage.New(config)
		if err != nil {
			return nil, err
		}
		err = stor.Connect(ctx)
		if err != nil {
			return nil, err
		}
		return stor, err
	default:
		return nil, errors.New("bad type store in config file")
	}
}

func (a *App) Close(ctx context.Context) error {
	a.servLog.Info("application Calendar stopped")

	err := a.storage.Close(ctx)
	if err != nil {
		a.servLog.Error("DB was closed with error:" + err.Error())
		return err
	}
	a.servLog.Info("application Calendar DB connection was closed")

	return nil
}

func (a *App) AddEvent(ctx context.Context, event *storage.Event, userID string) (string, error) {
	eventID, err := storage.UUID()
	if err != nil {
		return "", err
	}
	event.ID = eventID
	return eventID, a.storage.AddEvent(ctx, event, userID)
}

func (a *App) ChangeEvent(ctx context.Context, event *storage.Event) error {
	return a.storage.ChangeEvent(ctx, event)
}

func (a *App) ChangeEventUser(ctx context.Context, event *storage.Event, userID string) error {
	ev, err := a.storage.GetEvent(ctx, event.ID)
	if err != nil {
		return err
	}
	if ev.UserID != userID {
		return storage.ErrWrongEventUser
	}
	return a.storage.ChangeEvent(ctx, event)
}

func (a *App) GetEvent(ctx context.Context, eventID string) (*storage.Event, error) {
	return a.storage.GetEvent(ctx, eventID)
}

func (a *App) GetEventUser(ctx context.Context, eventID string, userID string) (*storage.Event, error) {
	event, err := a.storage.GetEvent(ctx, eventID)
	if err != nil {
		return nil, err
	}
	if event.UserID != userID {
		return nil, storage.ErrWrongEventUser
	}

	return a.storage.GetEvent(ctx, eventID)
}

func (a *App) DeleteEvent(ctx context.Context, eventID string) error {
	return a.storage.DeleteEvent(ctx, eventID)
}

func (a *App) DeleteEventUser(ctx context.Context, eventID string, userID string) error {
	event, err := a.storage.GetEvent(ctx, eventID)
	if err != nil {
		return err
	}
	if event.UserID != userID {
		return storage.ErrWrongEventUser
	}
	return a.storage.DeleteEvent(ctx, eventID)
}

func (a *App) ListEventsUser(ctx context.Context, from, to time.Time, userID string) ([]*storage.Event, error) {
	return a.storage.ListEventsUser(ctx, from, to, userID)
}

func (a *App) ListEvents(ctx context.Context, from, to time.Time) ([]*storage.Event, error) {
	return a.storage.ListEvents(ctx, from, to)
}

func (a *App) ListEventsDay(ctx context.Context, date time.Time, userID string) ([]*storage.Event, error) {
	from, to := datesDay(date)
	return a.storage.ListEventsUser(ctx, from, to, userID)
}

func (a *App) ListEventsWeek(ctx context.Context, date time.Time, userID string) ([]*storage.Event, error) {
	from, to := datesWeek(date)
	return a.storage.ListEventsUser(ctx, from, to, userID)
}

func (a *App) ListEventsMonth(ctx context.Context, date time.Time, userID string) ([]*storage.Event, error) {
	from, to := datesMonth(date)
	return a.storage.ListEventsUser(ctx, from, to, userID)
}

func (a *App) GetEventsToDelete(ctx context.Context, onTime time.Time) ([]*storage.Event, error) {
	return a.storage.ListEvents(ctx, storage.FistTimeCheckNotice, onTime.Add(-time.Hour*24*365))
}

func (a *App) ListNoticesToSend(ctx context.Context, onTime time.Time) ([]*storage.Notice, error) {
	return a.storage.ListNoticesToSend(ctx, onTime)
}

func datesDay(t time.Time) (from time.Time, to time.Time) {
	from = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	to = time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, time.UTC)
	return
}

func datesWeek(t time.Time) (from time.Time, to time.Time) {
	day := t.Day()
	if t.Weekday() == time.Sunday {
		day -= 6
	} else {
		day -= int(t.Weekday()) + 1
	}
	from = time.Date(t.Year(), t.Month(), day, 0, 0, 0, 0, time.UTC)
	to = time.Date(t.Year(), t.Month(), day+7, 0, 0, 0, 0, time.UTC)
	return
}

func datesMonth(t time.Time) (from time.Time, to time.Time) {
	from = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	to = time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	return
}
