package storage

import (
	"context"
	"errors"
	"time"
)

var (
	ErrEventIDNotFound = errors.New("event id not found")
	ErrWrongEventUser  = errors.New("event  doesn't belong to user")

	FistTimeCheckNotice = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
)

type Store interface {
	AddEvent(ctx context.Context, event *Event, userID string) error
	ChangeEvent(ctx context.Context, event *Event) error
	DeleteEvent(ctx context.Context, eventID string) error
	GetEvent(ctx context.Context, eventID string) (*Event, error)
	ListEvents(ctx context.Context, from time.Time, to time.Time, userID string) ([]*Event, error)
	ListNoticesToSend(ctx context.Context, onTime time.Time) ([]*Notice, error)
	GetLastNoticeTimeSetNew(ctx context.Context, lastCheck time.Time) (*time.Time, error)
	CountEvents(ctx context.Context, userID string) (int, error)
	Close(ctx context.Context) error
}
