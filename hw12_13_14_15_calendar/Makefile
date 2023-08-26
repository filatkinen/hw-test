BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar
	go build -v -o ./bin/calendar_scheduler  ./cmd/scheduler
	go build -v -o ./bin/calendar_sender  ./cmd/sender

run: build
	$(BIN) -config ./configs/server.yaml

build-img:
	docker-compose -f tests/integration-tests/docker-compose.yaml build

run-img: build-img
	docker-compose -f tests/integration-tests/docker-compose.yaml up -d


up: run-img

down:
	docker-compose -f tests/integration-tests/docker-compose.yaml down \
		 --rmi local \
		--volumes \
		--remove-orphans \
		--timeout 60; \


integration-tests: run-img
	test_status_code=0 ;\
	docker exec -it calendar-integration go test -v --count=1  ./tests/integration-tests/tests/ -tags integration || test_status_code=$$? ;\
	docker-compose -f tests/integration-tests/docker-compose.yaml down \
	   --rmi local \
       --volumes \
       --remove-orphans \
       --timeout 60 ;\
	exit $$test_status_code ;


version: build
	$(BIN) version

test:
	go test -race -v ./...

test-mysql:
	go test -v -race  ./internal/storage/ -tags mysql

test-pgsql:
	go test -v -race  ./internal/storage/ -tags pgsql

generate:
	go generate cmd/calendar/main.go

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.52.2

lint: install-lint-deps
	golangci-lint run ./...

migrate-mysql:
	./migrations/migrate-mysql.sh


migrate-pgsql:
	./migrations/migrate-pgsql.sh


.PHONY: build run build-img run-img version test lint
