FROM golang:1.20-alpine
WORKDIR /usr/src/app/

COPY . /usr/src/app/
COPY ./deployments/lab14/generator/calendar_config.yaml /usr/src/app/configs/calendar_config.yaml
RUN go build -v -o ./bin/generator  ./tests/generator-events


ENTRYPOINT [ "./bin/generator"]
CMD ["-config", "./configs/calendar_config.yaml"]
