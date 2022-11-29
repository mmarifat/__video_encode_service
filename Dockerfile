# syntax=docker/dockerfile:1

FROM golang:1.19.3-alpine

COPY go.mod ./
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod vendor
RUN go mod tidy
RUN apk add ffmpeg

COPY *.go ./

RUN swag init
RUN go build

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
