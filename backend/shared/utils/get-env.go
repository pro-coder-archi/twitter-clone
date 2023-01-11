package utils

import (
	"log"
	"os"
)

func GetEnv(envName string) string {
	envValue, isEnvFound := os.LookupEnv(envName)

	if(!isEnvFound) {
		log.Fatalf("❌ env %s not found", envName) }

	return envValue
}