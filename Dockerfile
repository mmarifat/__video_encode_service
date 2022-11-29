FROM golang:1.19.3-alpine

ENV GO111MODULE=on
ENV PORT=7000
WORKDIR /go/src/app
COPY go.mod .

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/githubnemo/CompileDaemon
RUN go mod vendor
RUN go mod tidy
RUN apk add ffmpeg
COPY . .
RUN swag init
RUN go build -o /docker-gs-ping
CMD [ "/docker-gs-ping" ]
