package config

import (
	"log"
	"os"
)

func CheckEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("WARNING!: Environment variable %s is missing", key)
	}
	return value
}
