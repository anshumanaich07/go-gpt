package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Localhost string
	Model     string
}

func LoadEnv() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("unable to load .env file: %v", err)
	}

	cfg := Config{}
	cfg.Localhost = os.Getenv("URL")
	cfg.Model = os.Getenv("MODEL")

	return &cfg, nil
}
