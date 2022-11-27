go mod vendor
go mod tidy
CompileDaemon -command="./video-conversion-service" -include=".env"