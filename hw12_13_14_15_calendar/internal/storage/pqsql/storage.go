package pgsqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
	_ "github.com/lib/pq" // import pq
)

type Storage struct { // TODO
	db       *sql.DB
	dsn      string
	dbconfig server.DBConfig
}

func New(config server.Config) (*Storage, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&Timezone=UTC",
		config.DB.DBUser, config.DB.DBPass, config.DB.DBAddress, config.DB.DBPort, config.DB.DBName)
	db, err := openDB(config.DB, dsn)
	if err != nil {
		return nil, err
	}
	return &Storage{
		db:       db,
		dsn:      dsn,
		dbconfig: config.DB,
	}, nil
}

func openDB(config server.DBConfig, dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxIdleTime(config.MaxIdleTime)
	return db, nil
}

func (s *Storage) Connect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.db.PingContext(ctx)
}

func (s *Storage) GetLastNoticeTimeSetNew(ctx context.Context, onTime time.Time) (lastCheck *time.Time, err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err = errors.Join(err, tx.Rollback())
		}
	}()
	query := `SELECT last_check_date_time
	FROM notes_check
	ORDER BY last_check_date_time DESC
	LIMIT 1`

	err = tx.QueryRowContext(ctx, query).Scan(lastCheck)

	if errors.Is(err, sql.ErrNoRows) {
		*lastCheck = storage.FistTimeCheckNotice
	} else if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, `TRUNCATE notes_check`)
	if err != nil {
		return nil, err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO notes_check (last_check_date_time) 
										VALUES ($1)`, s.truncateTime(onTime))
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return
}

func (s *Storage) Close(_ context.Context) error {
	return s.db.Close()
}

func (s *Storage) AddEvent(ctx context.Context, event *storage.Event, userID string) error {
	query := `INSERT INTO events (event_id, title, description, 
                    date_time_start, date_time_end, date_time_notice, user_id) 
			  VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := s.db.ExecContext(ctx, query, event.ID, event.Title, event.Description,
		event.DateTimeStart, event.DateTimeEnd, event.DateTimeNotice, userID)
	return err
}

func (s *Storage) GetEvent(ctx context.Context, eventID string) (*storage.Event, error) {
	r := storage.Event{}
	var description sql.NullString
	query := `SELECT event_id, title, description, 
                    date_time_start, date_time_end, date_time_notice, user_id from events WHERE event_id=$1`
	if err := s.db.QueryRowContext(ctx, query, eventID).
		Scan(&r.ID, &r.Title, &description, &r.DateTimeStart, &r.DateTimeEnd, &r.DateTimeNotice, &r.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrEventIDNotFound
		}
		return nil, err
	}
	r.Description = description.String
	return &r, nil
}

func (s *Storage) ChangeEvent(ctx context.Context, event *storage.Event) error {
	query := `UPDATE events SET 
                title=$1,
				description=$2,
				date_time_start=$3,
				date_time_end=$4,
				date_time_notice=$5,
				user_id=$6 
              WHERE event_id=$7`
	args := []any{
		event.Title,
		event.Description,
		event.DateTimeStart,
		event.DateTimeEnd,
		event.DateTimeNotice,
		event.UserID,
		event.ID,
	}
	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return storage.ErrEventIDNotFound
	}
	return nil
}

func (s *Storage) DeleteEvent(ctx context.Context, eventID string) error {
	query := `DELETE FROM events 
              WHERE event_id=$1`
	result, err := s.db.ExecContext(ctx, query, eventID)
	if err != nil {
		return err
	}
	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return storage.ErrEventIDNotFound
	}
	return nil
}

func (s *Storage) ListEventsUser(ctx context.Context, from, to time.Time, userID string) ([]*storage.Event, error) {
	query := `SELECT event_id, title, description, 
                    date_time_start, date_time_end, date_time_notice, user_id 
			  FROM events
			  WHERE user_id=$1 AND date_time_start>=$2 AND date_time_start<=$3`
	rows, err := s.db.QueryContext(ctx, query, userID, s.truncateTime(from), s.truncateTime(to))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*storage.Event
	for rows.Next() {
		var event storage.Event
		var description sql.NullString
		err := rows.Scan(
			&event.ID,
			&event.Title,
			&description,
			&event.DateTimeStart,
			&event.DateTimeEnd,
			&event.DateTimeNotice,
			&event.UserID,
		)
		if err != nil {
			return nil, err
		}
		event.Description = description.String
		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Storage) ListEvents(ctx context.Context, from, to time.Time) ([]*storage.Event, error) {
	query := `SELECT event_id, title, description, 
                    date_time_start, date_time_end, date_time_notice, user_id 
			  FROM events
			  WHERE date_time_start>=$1 AND date_time_start<=$2`
	rows, err := s.db.QueryContext(ctx, query, s.truncateTime(from), s.truncateTime(to))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*storage.Event
	for rows.Next() {
		var event storage.Event
		var description sql.NullString
		err := rows.Scan(
			&event.ID,
			&event.Title,
			&description,
			&event.DateTimeStart,
			&event.DateTimeEnd,
			&event.DateTimeNotice,
			&event.UserID,
		)
		if err != nil {
			return nil, err
		}
		event.Description = description.String
		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Storage) ListNoticesToSend(ctx context.Context, onTime time.Time) ([]*storage.Notice, error) {
	dateTimeNoticeLast, err := s.GetLastNoticeTimeSetNew(ctx, onTime)
	if err != nil {
		return nil, err
	}
	query := `SELECT event_id, title, date_time_start, user_id 
			  FROM events
			  WHERE 
			    (date_time_notice<=$1  
			    AND date_time_start>$2 
			    AND date_time_notice>$3)`
	rows, err := s.db.QueryContext(ctx, query,
		s.truncateTime(onTime), s.truncateTime(onTime), s.truncateTime(*dateTimeNoticeLast))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notices []*storage.Notice
	for rows.Next() {
		var event storage.Event
		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.DateTimeStart,
			&event.UserID,
		)
		if err != nil {
			return nil, err
		}
		notices = append(notices, &storage.Notice{
			ID:       event.ID,
			Title:    event.Title,
			DateTime: event.DateTimeStart,
			UserID:   event.UserID,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notices, nil
}

func (s *Storage) CountEvents(ctx context.Context, userID string) (int, error) {
	query := `SELECT COUNT(*) AS count FROM events WHERE user_id=$1`
	var count int
	err := s.db.QueryRowContext(ctx, query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *Storage) truncateTime(t time.Time) time.Time {
	return t.Round(time.Microsecond)
}
