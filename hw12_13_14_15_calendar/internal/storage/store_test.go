package storage_test

import (
	"context"
	"testing"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/stretchr/testify/require"
)

var (
	timeNow = time.Now().UTC().Round(0).Add(-time.Hour * 24 * 365)
	UserID  = "ba2881ed-694d-46e7-b095-1372bb79e14f"
	events  = []*storage.Event{
		{
			ID:             "ba2881ed-694d-46e7-b095-1372bb79e14f",
			Title:          "Task1",
			DateTimeStart:  timeNow.Add(time.Second * 10),
			DateTimeEnd:    timeNow.Add(time.Second * 30),
			DateTimeNotice: timeNow.Add(-time.Second * 10),
			UserID:         UserID,
		},
		{
			ID:             "40d42cfa-6a3b-5544-f253-a21e899f0bc9",
			Title:          "Task2",
			DateTimeStart:  timeNow.Add(time.Second * 10),
			DateTimeEnd:    timeNow.Add(time.Second * 30),
			DateTimeNotice: timeNow.Add(-time.Second * 10),
			UserID:         UserID,
		},
		{
			ID:             "5eb29ff9-ea25-4894-a6fd-052ca84531b0",
			Title:          "Task3",
			DateTimeStart:  timeNow.Add(time.Second * 10),
			DateTimeEnd:    timeNow.Add(time.Second * 30),
			DateTimeNotice: timeNow.Add(-time.Second * 5),
			UserID:         UserID,
		},
		{
			ID:             "aaa1d209-3181-15c9-4dbe-400b80331a78",
			Title:          "Task4",
			DateTimeStart:  timeNow.Add(time.Second * 10),
			DateTimeEnd:    timeNow.Add(time.Second * 30),
			DateTimeNotice: timeNow.Add(time.Second * 0),
			UserID:         UserID,
		},
		{
			ID:             "732f2d0d-2630-0cd5-f78a-732faf83e68a",
			Title:          "Task5",
			DateTimeStart:  timeNow.Add(time.Second * 10),
			DateTimeEnd:    timeNow.Add(time.Second * 30),
			DateTimeNotice: timeNow.Add(time.Second * 5),
			UserID:         UserID,
		},
	}
)

func TestMemoryStorage(t *testing.T) {
	ctx := context.Background()
	defer ctx.Done()
	memoryStorage := memorystorage.New()
	var s storage.Store = memoryStorage
	runTestStorage(t, ctx, s)
}

func runTestStorage(t *testing.T, ctx context.Context, store storage.Store) { //nolint
	// delete events probably added by test earlier
	for i := range events {
		_ = store.DeleteEvent(ctx, events[i].ID)
	}

	// added events for testing
	for i := range events {
		_ = store.AddEvent(ctx, events[i], UserID)
	}

	// running tests
	t.Run("Count rows", func(t *testing.T) {
		count, err := store.CountEvents(ctx, UserID)
		require.Nil(t, err)
		require.True(t, count > 0)
	})

	t.Run("list Events", func(t *testing.T) {
		idevents := make(map[string]bool, len(events))
		for _, v := range events {
			idevents[v.ID] = true
		}

		ev, err := store.ListEvents(ctx, timeNow, timeNow.Add(time.Second*30), UserID)
		require.Nil(t, err)
		require.NotNil(t, ev)
		for i := range events {
			_, ok := idevents[ev[i].ID]
			require.True(t, ok)
		}
	})

	t.Run("Delete Event num 3", func(t *testing.T) {
		_, err := store.GetEvent(ctx, events[2].ID)
		require.Nil(t, err)
		err = store.DeleteEvent(ctx, events[2].ID)
		require.Nil(t, err)
		_, err = store.GetEvent(ctx, events[2].ID)
		require.Equal(t, err, storage.ErrEventIDNotFound)
	})
	t.Run("Change Event 2", func(t *testing.T) {
		chEvent := events[1]
		chEvent.Title = "Task2 new name"
		_ = store.ChangeEvent(ctx, chEvent)
		newEvent, err := store.GetEvent(ctx, chEvent.ID)
		require.Nil(t, err)
		require.Equal(t, newEvent.ID, chEvent.ID)
		require.Equal(t, newEvent.Title, chEvent.Title)
	})
	t.Run("Get ListNoticesToSend", func(t *testing.T) {
		// saving time point of checking notice and putting there FistTimeCheckNotice value
		lastTimeNoticesCheck, err := store.GetLastNoticeTimeSetNew(ctx, storage.FistTimeCheckNotice)
		require.Nil(t, err)

		// we want to get 2 notices
		timeToGetNotice := timeNow.Add(-time.Second * 5)
		notices, err := store.ListNoticesToSend(ctx, timeToGetNotice)
		require.Nil(t, err)
		require.Equal(t, 2, len(notices))

		// we want to get next 1 notice
		timeToGetNotice = timeNow.Add(time.Second * 1)
		notices, err = store.ListNoticesToSend(ctx, timeToGetNotice)
		require.Nil(t, err)
		require.Equal(t, 1, len(notices))

		// we want to get next 1 notice
		timeToGetNotice = timeNow.Add(+time.Second * 6)
		notices, err = store.ListNoticesToSend(ctx, timeToGetNotice)
		require.Nil(t, err)
		require.Equal(t, 1, len(notices))

		// We want to get 0 notes
		timeToGetNotice = timeNow.Add(+time.Second * 8)
		notices, err = store.ListNoticesToSend(ctx, timeToGetNotice)
		require.Nil(t, err)
		require.Nil(t, notices)

		// restore time point
		_, err = store.GetLastNoticeTimeSetNew(ctx, *lastTimeNoticesCheck)
		require.Nil(t, err)
	})

	// redo - delete events added in test earlier: our tests could fatal after adding testing event's records
	for i := range events {
		_ = store.DeleteEvent(ctx, events[i].ID)
	}
}
