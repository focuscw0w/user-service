package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SecretKey string
	Port      string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		SecretKey: os.Getenv("SECRET_KEY"),
		Port:      os.Getenv("PORT"),
	}

	return cfg, nil
}
