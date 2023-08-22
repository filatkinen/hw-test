package memorystorage

import (
	"context"
	"sync"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	events              []*storage.Event
	lastTimeCheckNotice time.Time
	mu                  sync.RWMutex
}

func New() *Storage {
	return &Storage{
		events:              []*storage.Event{},
		lastTimeCheckNotice: storage.FistTimeCheckNotice,
		mu:                  sync.RWMutex{},
	}
}

func (s *Storage) AddEvent(_ context.Context, event *storage.Event, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	ev := *event
	ev.UserID = userID
	s.events = append(s.events, &ev)
	return nil
}

func (s *Storage) GetEvent(_ context.Context, eventID string) (*storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.events {
		if s.events[i].ID == eventID {
			ev := *s.events[i]
			return &ev, nil
		}
	}
	return nil, storage.ErrEventIDNotFound
}

func (s *Storage) ChangeEvent(_ context.Context, event *storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.events {
		if s.events[i].ID == event.ID {
			*s.events[i] = *event
			return nil
		}
	}
	return storage.ErrEventIDNotFound
}

func (s *Storage) DeleteEvent(_ context.Context, eventID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.events {
		if s.events[i].ID == eventID {
			s.events = append(s.events[:i], s.events[i+1:]...)
			return nil
		}
	}
	return storage.ErrEventIDNotFound
}

func (s *Storage) ListEventsUser(_ context.Context, from, to time.Time, userID string) ([]*storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var events []*storage.Event
	for i := range s.events {
		if s.events[i].UserID == userID &&
			(s.events[i].DateTimeStart.Equal(from) || s.events[i].DateTimeStart.Equal(to) ||
				(s.events[i].DateTimeStart.After(from) && s.events[i].DateTimeStart.Before(to))) {
			events = append(events, s.events[i])
		}
	}
	return events, nil
}

func (s *Storage) ListEvents(_ context.Context, from, to time.Time) ([]*storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var events []*storage.Event
	for i := range s.events {
		if s.events[i].DateTimeStart.Equal(from) || s.events[i].DateTimeStart.Equal(to) ||
			(s.events[i].DateTimeStart.After(from) && s.events[i].DateTimeStart.Before(to)) {
			events = append(events, s.events[i])
		}
	}
	return events, nil
}

func (s *Storage) ListNoticesToSend(_ context.Context, onTime time.Time) ([]*storage.Notice, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var notice []*storage.Notice
	for i := range s.events {
		if s.events[i].DateTimeNotice.Equal(onTime) ||
			s.events[i].DateTimeNotice.Before(onTime) && s.events[i].DateTimeStart.After(onTime) &&
				s.events[i].DateTimeNotice.After(s.lastTimeCheckNotice) {
			notice = append(notice, &storage.Notice{
				ID:       s.events[i].ID,
				Title:    s.events[i].Title,
				DateTime: s.events[i].DateTimeStart,
				UserID:   s.events[i].UserID,
			})
		}
	}
	s.lastTimeCheckNotice = onTime
	return notice, nil
}

func (s *Storage) DeleteOldEvents(_ context.Context, onTime time.Time) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var count int
	for i := range s.events {
		if s.events[i].DateTimeStart.Before(onTime) {
			s.events = append(s.events[:i], s.events[i+1:]...)
			count++
		}
	}
	return count, nil
}

func (s *Storage) CountEvents(_ context.Context, userID string) (int, error) {
	count := 0
	for i := range s.events {
		if s.events[i].UserID == userID {
			count++
		}
	}
	return count, nil
}

func (s *Storage) GetLastNoticeTimeSetNew(_ context.Context, onTime time.Time) (*time.Time, error) {
	lastTime := s.lastTimeCheckNotice
	s.lastTimeCheckNotice = onTime
	return &lastTime, nil
}

func (s *Storage) Close(_ context.Context) error {
	return nil
}
