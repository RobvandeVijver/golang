package config

import (
	"os"
)

func GetHost() string {
	host := os.Getenv("API_HOST")
	if host == "" {
		host = "localhost:8090"
	}
	return host
}
