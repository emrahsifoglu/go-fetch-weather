FROM golang:1.21.0-alpine3.18

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o go-fetch-weather ./cmd

CMD ["/app/go-fetch-weather"]
