FROM golang:1.19.3-alpine

ENV GO111MODULE=on
ENV PORT=7000
WORKDIR /go/src/app
COPY go.mod /go/src/app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod vendor
RUN go mod tidy
RUN apk add ffmpeg
COPY . /go/src/app
RUN swag init
RUN go build -o /go/src/app/docker-gs-ping
CMD [ "/go/src/app/docker-gs-ping" ]
