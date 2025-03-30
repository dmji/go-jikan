package utils

import (
	"os"
)

// GetEnv returns the value of the environment variable named by the key or the defaultValue if the environment variable is not set.
func GetEnv(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}
