package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}

	return os.Getenv(key)
}
