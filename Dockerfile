# GO Repo base repo
FROM golang:1.19.3-alpine as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod ./

# Download all the dependencies
RUN go install github.com/cosmtrek/air@latest
RUN go mod vendor
RUN go mod tidy

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# GO Repo base repo
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl ffmpeg

RUN mkdir /app

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8000
EXPOSE 8000

# Run Executable
CMD ["./main"]
