# syntax=docker/dockerfile:1

FROM golang:1.18.2-alpine

WORKDIR /app

COPY go.mod ./
RUN go install github.com/cosmtrek/air@latest
RUN go mod vendor
RUN go mod tidy
RUN apk add ffmpeg

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
