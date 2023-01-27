FROM golang:1.20rc3-alpine3.17

ADD ./src /app

WORKDIR /app

RUN go build .

ENV APIKEY "AIzaSyB-3nxRIQgI8CAsL8u1MXh_HIFbKwa9xWg"

CMD ["./tyvamcbe"]
