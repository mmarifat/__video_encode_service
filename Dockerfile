# syntax=docker/dockerfile:1
FROM golang:1.19.3-alpine as builder

WORKDIR /app

COPY go.mod ./

# Download all the dependencies
RUN go install github.com/cosmtrek/air@latest
RUN go mod vendor
RUN go mod tidy

COPY *.go ./

RUN go build -o /docker-gs-ping
