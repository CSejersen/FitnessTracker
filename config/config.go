package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	JWTSecret string
	Port      string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using system environment variables")
	}

	config := &Config{
		JWTSecret: getEnv("JWT_SECRET", ""),
		Port:      getEnv("PORT", ":8080"),
	}

	return config, nil
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
