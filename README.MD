VIDEO-CONVERSION-SERVICE

## Getting Started

Install dependencies

```bash
$ go mod vendor
# and
$ go mod tidy
```

## Environment

To edit environment variables, create a file with name `.env` and copy the contents from `.env.example` to start with.

Run Locally

```bash
$ ./start-dev.sh
```

Build for production

```bash
$ swag init
# then
$ go build
```

## Logging

The application uses default logger of gin. The configuration file is at `main.go`.

* All logs are saved in `./logs` directory.
* Success logs save inside `.logs/success.log`
* Error logs save inside `.logs/error.log`