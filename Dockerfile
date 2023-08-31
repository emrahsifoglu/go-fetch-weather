FROM golang:1.21.0-alpine3.18

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download
RUN go build -o go-fetch-weather ./cmd

EXPOSE 8081

CMD ["/app/go-fetch-weather"]
