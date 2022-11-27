go mod vendor
go mod tidy
swag init
CompileDaemon -command="./video-conversion-service" -include=".env"