package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func NewConfig() (*Config, error) {
	//loading .env file with configuration variables
	_ = godotenv.Load()

	//getting port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//getting URL
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is required")
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}, nil

}

func NewConfigMust() *Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load core config: %v", err))
	}
	return cfg
}
