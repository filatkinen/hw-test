CREATE TABLE IF NOT EXISTS events
(
    event_id         CHAR(36)     NOT NULL
        PRIMARY KEY,
    title            VARCHAR(255) NOT NULL,
    description      TEXT         NULL,
    date_time_start  DATETIME     NOT NULL,
    date_time_end    DATETIME     NOT NULL,
    date_time_notice DATETIME     NOT NULL,
    user_id          CHAR(36)     NOT NULL
)
    COLLATE = utf8mb4_general_ci;

CREATE INDEX events_date_time_notice_date_time_start_index
    ON events (date_time_notice, date_time_start);

CREATE INDEX events_date_time_start_index
    ON events (date_time_start);

CREATE TABLE IF NOT EXISTS notes_check
(
    last_check_date_time DATETIME NOT NULL
);