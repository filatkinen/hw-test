#!/usr/bin/bash

CALENDAR_DB_USER='calendar'
CALENDAR_DB_PASS='pass'

export DSN_CALENDAR_PQSQL='postgresql://'$CALENDAR_DB_USER:$CALENDAR_DB_PASS'@localhost:5432/calendar'

migrate -database  $DSN_CALENDAR_PQSQL  -path ./migrations/pgsql -verbose up
