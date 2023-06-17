#!/usr/bin/bash

CALENDAR_DB_USER='calendar'
CALENDAR_DB_PASS='pass'

DSN_CALENDAR_MYSQL='mysql://'$CALENDAR_DB_USER:$CALENDAR_DB_PASS'@tcp(localhost:3306)/calendar'

migrate  -database  $DSN_CALENDAR_MYSQL -path ./migrations/mysql  up

