package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// LoadEnv loads environment variables from a .env file if it exists.
// Logs a warning if the file is missing but continues with OS environment variables.
func LoadEnv(logger *zap.Logger) {
	if err := godotenv.Load(); err != nil {
		if logger != nil {
			logger.Warn("No .env file found, relying on environment variables")
		} else {
			// fallback to standard log if logger is not initialized yet
			println("Warning: No .env file found, relying on environment variables")
		}
	}
}

// GetEnv returns the value of the environment variable if set, otherwise returns the fallback.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
