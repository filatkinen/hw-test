//go:build mysql

package storage_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
	mysqlstorage "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage/mysql"
)

func TestMysqlStorage(t *testing.T) {
	_ = os.Setenv("CALENDAR_DB_USER", "calendar")
	_ = os.Setenv("CALENDAR_DB_PASS", "pass")
	conf, err := server.NewConfig("../../configs/calendar_config.yaml")
	if err != nil {
		log.Fatalf("error reading config file %v", err)
	}
	mysqlStorage, err := mysqlstorage.New(conf)
	if err != nil {
		log.Fatalf("%v", err)
	}

	ctx := context.Background()
	defer ctx.Done()

	err = mysqlStorage.Connect(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer mysqlStorage.Close(ctx)

	var s storage.Store = mysqlStorage
	runTestStorage(t, ctx, s)
}
