FROM golang:1.19-alpine

ENV GO111MODULE=on

RUN apk add ffmpeg

WORKDIR /video-conversion-service

COPY go.mod ./
COPY go.sum ./

RUN go mod vendor
RUN go mod download
RUN go mod tidy

COPY vendor /video-conversion-service
COPY . /video-conversion-service

RUN go build -o /video-conversion-service

ENV PORT=9595

CMD ["/video-conversion-service/video-conversion-service"]
