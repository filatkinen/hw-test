//go:build pgsql

package storage_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
	pgsqlstorage "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage/pqsql"
)

func TestPgsqlStorage(t *testing.T) {
	_ = os.Setenv("CALENDAR_DB_USER", "calendar")
	_ = os.Setenv("CALENDAR_DB_PASS", "pass")
	conf, err := server.NewConfig("../../configs/calendar_config.yaml")
	if err != nil {
		log.Fatalf("error reading config file %v", err)
	}
	conf.DB.DBPort = "5432"
	pgsqlStorage, err := pgsqlstorage.New(conf)
	if err != nil {
		log.Fatalf("%v", err)
	}

	ctx := context.Background()
	defer ctx.Done()

	err = pgsqlStorage.Connect(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer pgsqlStorage.Close(ctx)

	var s storage.Store = pgsqlStorage
	runTestStorage(t, ctx, s)
}
