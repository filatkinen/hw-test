FROM golang:1.20-alpine
WORKDIR /usr/src/app/

COPY . /usr/src/app/
COPY ./deployments/lab14/calendar/calendar_config.yaml /usr/src/app/configs/calendar_config.yaml
RUN go build -v -o ./bin/calendar  ./cmd/calendar


ENTRYPOINT [ "./bin/calendar"]
CMD ["-config", "./configs/calendar_config.yaml"]
EXPOSE 8800
