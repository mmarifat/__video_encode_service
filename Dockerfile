FROM golang:1.19-alpine as builder

ENV GO111MODULE=on

RUN apk add ffmpeg

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
# Expose port 8080 to the outside world
EXPOSE 9595
# Command to run the executable
CMD ["./main"]
