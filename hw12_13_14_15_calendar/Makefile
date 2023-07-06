BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

run: build
	$(BIN) -config ./configs/server.yaml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race -v ./...

test-mysql:
	go test -v -race  ./internal/storage/ -tags mysql

test-pgsql:
	go test -v -race  ./internal/storage/ -tags pgsql

generate:
	protoc -I  internal/grpcservice/ internal/grpcservice/calendar.proto   --go_out=internal/grpcservice/
	protoc -I  internal/grpcservice/ internal/grpcservice/calendar.proto   --go-grpc_out=require_unimplemented_servers=false:internal/grpcservice/



install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.52.2

lint: install-lint-deps
	golangci-lint run ./...

migrate-mysql:
	./migrations/migrate-mysql.sh


migrate-pgsql:
	./migrations/migrate-pgsql.sh


.PHONY: build run build-img run-img version test lint
