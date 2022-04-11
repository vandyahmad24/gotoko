package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}
