package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDSN string
	Port        string
}

func LoadConfig() *Config {

	godotenv.Load()

	return &Config{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
		Port:        os.Getenv("PORT"),
	}
}