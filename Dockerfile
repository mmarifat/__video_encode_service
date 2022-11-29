FROM golang:latest

ENV GO111MODULE=on
ENV PORT=9595

WORKDIR /app
COPY go.mod /app
COPY go.sum /app

RUN go install -mod=mod github.com/swaggo/swag/cmd/swag
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN go mod vendor
RUN go mod download
RUN go mod tidy
COPY . /app
ENTRYPOINT CompileDaemon --build="go build -o video-conversion-service" --command=./video-conversion-service
