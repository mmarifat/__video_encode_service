VIDEO-CONVERSION-SERVICE

## Requirements

- go (version 1.19+)
- air
- ffmpeg
- redis

## Getting Started

Install dependencies

```bash
$ go install github.com/cosmtrek/air@latest
# then
$ go mod vendor
# and
$ go mod tidy
```

## Environment

To edit environment variables, create a file with name `.env` and copy the contents from `.env.example` to start with.

Run Locally

```bash
$ air
```

Build for production

```bash
$ swag init
# then
$ go build -o video-conversion-service
# then
$ ./video-conversion-service
```

## Logging

The application uses default logger of gin. The configuration file is at `main.go`.

* All logs are saved in `./logs` directory.
* Success logs save inside `.logs/success.log`
* Error logs save inside `.logs/error.log`