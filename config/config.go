package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold the configuration values
type Config struct {
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	AWSRegion          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	ServerPort         string
}

// LoadConfig function loads environment variables from the .env file
func LoadConfig() *Config {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Return the configuration values
	return &Config{
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		AWSRegion:          os.Getenv("AWS_REGION"),
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		ServerPort:         os.Getenv("SERVER_PORT"),
	}
}
