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
	servLog.Logging(logger.LevelInfo, "application Calendar started")
	servLog.Logging(logger.LevelInfo, "application Calendar using db:"+config.StoreType)

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
	a.servLog.Logging(logger.LevelInfo, "application Calendar stopped")

	err := a.storage.Close(ctx)
	if err != nil {
		a.servLog.Error("DB was closed with error:" + err.Error())
		return err
	}
	a.servLog.Logging(logger.LevelInfo, "application Calendar DB connection was closed")

	return nil
}

func (a *App) CreateEvent(ctx context.Context, event *storage.Event) error {
	return a.storage.AddEvent(ctx, event)
}

func (a *App) ListEvents(ctx context.Context, from, to time.Time) ([]*storage.Event, error) {
	return a.storage.ListEvents(ctx, from, to)
}
