package storage

import (
	"context"
	"errors"
	"time"
)

var (
	ErrEventIDNotFound  = errors.New("event id not found")
	FistTimeCheckNotice = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
)

type Store interface {
	AddEvent(context.Context, *Event) error
	GetEvent(context.Context, string) (*Event, error)
	ChangeEvent(context.Context, *Event) error
	DeleteEvent(context.Context, string) error
	ListEvents(context.Context, time.Time, time.Time) ([]*Event, error)
	ListNoticesToSend(context.Context, time.Time) ([]*Notice, error)
	GetLastNoticeTimeSetNew(context.Context, time.Time) (*time.Time, error)
	CountEvents(context.Context) (int, error)
	Close(context.Context) error
}
