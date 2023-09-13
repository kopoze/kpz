package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Get value from `.env` file.
func GetEnv(key string) string {
	Err := godotenv.Load(".env")
	if Err != nil {
		log.Fatal("Error loading .env file", Err)
	}
	return os.Getenv(key)
}
