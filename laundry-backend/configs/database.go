package configs

import (
	"os"
)

var (
	DB_HOST     = getEnv("DB_HOST", "localhost")
	DB_PORT     = getEnv("DB_PORT", "5432")
	DB_USER     = getEnv("DB_USER", "user")
	DB_PASSWORD = getEnv("DB_PASSWORD", "password")
	DB_NAME     = getEnv("DB_NAME", "laundry")
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}