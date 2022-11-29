FROM golang:1.19.3-alpine

ENV GO111MODULE=on
ENV PORT=7000
WORKDIR /app
COPY go.mod /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod vendor
RUN go mod tidy
RUN apk add ffmpeg
COPY . /app
RUN swag init
RUN go build -o mai
ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main
