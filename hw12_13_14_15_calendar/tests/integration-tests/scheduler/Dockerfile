FROM golang:1.20-alpine
WORKDIR /usr/src/app/

COPY . /usr/src/app/
COPY ./tests/integration-tests/scheduler/scheduler_config.yaml /usr/src/app/configs/scheduler_config.yaml
RUN go build -v -o ./bin/scheduler  ./cmd/scheduler


ENTRYPOINT [ "./bin/scheduler"]
CMD ["-config", "./configs/scheduler_config.yaml"]
