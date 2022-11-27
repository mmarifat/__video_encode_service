package funtions

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file;" + err.Error())
	}
	return os.Getenv(key)
}
