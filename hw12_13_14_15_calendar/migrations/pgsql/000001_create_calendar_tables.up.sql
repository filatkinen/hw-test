CREATE TABLE IF NOT EXISTS events
(
    event_id         CHAR(36)                 NOT NULL
        CONSTRAINT events_pkey
        PRIMARY KEY,
    title            VARCHAR(255)             NOT NULL,
    description      TEXT,
    date_time_start  TIMESTAMP WITH TIME ZONE NOT NULL,
    date_time_end    TIMESTAMP WITH TIME ZONE NOT NULL,
    date_time_notice TIMESTAMP WITH TIME ZONE NOT NULL,
    user_id          CHAR(36)                 NOT NULL
);

ALTER TABLE events
    OWNER TO calendar;

CREATE INDEX IF NOT EXISTS events_date_time_notice_date_time_start_index
    ON events (date_time_notice, date_time_start);

CREATE INDEX IF NOT EXISTS events_date_time_start_index
    ON events (date_time_start);

CREATE TABLE IF NOT EXISTS notes_check
(
    last_check_date_time TIMESTAMP WITH TIME ZONE NOT NULL
);

