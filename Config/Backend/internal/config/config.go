package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	Port     string
	DBName   string
}

func LoadConfig() *Config {

	godotenv.Load()

	return &Config{
		MongoURI: os.Getenv("MONGODB_URI"),
		Port:     os.Getenv("PORT"),
		DBName:   os.Getenv("DBNAME"),
	}
}