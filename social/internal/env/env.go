package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func GetInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		parsedValue, err := strconv.Atoi(value)
		if err == nil {
			return parsedValue
		}

		panic(fmt.Errorf("failed to parse %s: %w", key, err))
	}
	return defaultValue
}
