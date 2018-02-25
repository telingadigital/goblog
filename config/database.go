package config

import (
	// "github.com/joho/godotenv"
	// "log"
	// "os"
)

// // Config stores configuration values
// type Config struct {
// 	DBhost string
// 	DBusername string
// 	DBdatabase string
// 	DBpassword string
// }

// // Set Default Configuration
// func setConfig() *Config {
// 	env := godotenv.Load()

// 	if env != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	return &Config{
// 		DBhost: os.Getenv("DB_HOST"),
// 		DBusername: os.Getenv("DB_USERNAME"),
// 		DBdatabase: os.Getenv("DB_DATABASE"),
// 		DBpassword: os.Getenv("DB_PASSWORD"),
// 	}
// }