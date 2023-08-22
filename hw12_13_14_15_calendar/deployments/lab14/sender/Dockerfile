FROM golang:1.20-alpine
WORKDIR /usr/src/app/

COPY . /usr/src/app/
COPY ./deployments/lab14/sender/sender_config.yaml /usr/src/app/configs/sender_config.yaml
RUN go build -v -o ./bin/sender  ./cmd/sender


ENTRYPOINT [ "./bin/sender"]
CMD ["-config", "./configs/sender_config.yaml"]
