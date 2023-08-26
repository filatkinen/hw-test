FROM golang:1.20-alpine
WORKDIR /usr/src/app/

COPY . /usr/src/app/
RUN go mod vendor

ENTRYPOINT ["tail"]
CMD ["-f","/dev/null"]
