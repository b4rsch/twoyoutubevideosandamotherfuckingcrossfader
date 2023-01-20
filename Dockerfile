FROM golang:1.20rc3-alpine3.17

ADD ./src /app

WORKDIR /app

RUN go build .

CMD ["./tyvamcbe"]
