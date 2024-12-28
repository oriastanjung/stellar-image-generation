package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	IMAGE_GENERATION_MODEL          string
	IMAGE_API_GENERATION_URL        string
	IMAGE_API_GENERATION_ORIGIN     string
	IMAGE_API_GENERATION_VALIDATED  string
	IMAGE_API_GENERATION_MESSAGE_ID string
}

func LoadEnv() *Config {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config := &Config{
		IMAGE_GENERATION_MODEL:          getEnv("IMAGE_GENERATION_MODEL", ""),
		IMAGE_API_GENERATION_URL:        getEnv("IMAGE_API_GENERATION_URL", ""),
		IMAGE_API_GENERATION_ORIGIN:     getEnv("IMAGE_API_GENERATION_ORIGIN", ""),
		IMAGE_API_GENERATION_VALIDATED:  getEnv("IMAGE_API_GENERATION_VALIDATED", ""),
		IMAGE_API_GENERATION_MESSAGE_ID: getEnv("IMAGE_API_GENERATION_MESSAGE_ID", ""),
	}

	return config
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not set, defaulting to %s", key, fallback)
		value = fallback
	}
	return value

}
