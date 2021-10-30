package env

import (
	"log"
	"os"
	"strconv"
)

func MustEnvString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("get env string key = %s,failed\n", key)
	}
	return value
}

func MustEnvInt(key string) int {
	castedInt, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Fatalf("get env int key = %s failed\n", key)
	}
	if castedInt == 0 {
		log.Fatalf("get env int key = %s failed, value is default\n", key)
	}
	return castedInt
}

