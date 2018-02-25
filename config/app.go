package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config stores configuration values
type Config struct {
	AppName string
	AppUrl string
	Port string
	Env string
}

// Set Default Configuration
func SetConfig() *Config {
	env := godotenv.Load()

	if env != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		AppName: os.Getenv("APP_NAME"),
		AppUrl: os.Getenv("APP_URL"),
		Port: os.Getenv("PORT"),
		Env: os.Getenv("APP_ENV"),
	}
}