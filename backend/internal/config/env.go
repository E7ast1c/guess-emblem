package config

import (
	"log"
	"os"
)

func MustEnvString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("get env string failed, key = %s\n", key)
	}
	return value
}
