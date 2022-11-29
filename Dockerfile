# syntax=docker/dockerfile:1

FROM golang:1.19.3-alpine

WORKDIR /app

COPY go.mod ./
RUN go install github.com/cosmtrek/air@latest
RUN go mod download github.com/swaggo/swag
RUN go install github.com/swaggo/swag/cmd/swag
RUN go mod vendor
RUN go mod tidy
RUN apk add ffmpeg

COPY *.go ./

RUN swag init
RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
