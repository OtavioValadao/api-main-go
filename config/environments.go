package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvironments(key string) string {
	logger = GetLogger("environment")
	godotenv.Load(".env")
	value := os.Getenv(key)
	if value == "" {
		logger.Infof("The environment %v is empty", key)
	}
	return value
}
