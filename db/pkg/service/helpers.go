package db

import (
	"os"

	"github.com/joho/godotenv"
)

// getDBName retrieves the database name from the environment variables.
func getDBName() string {
	return os.Getenv("DB_NAME")
}

// getDBUrl retrieves the database url from the environment variables.
func getDBUrl() string {
	return os.Getenv("DB_URL")
}

// loadEnv loads the environment variables from the .env file.
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

// Initialize the environment variables.
func init() {
	err := loadEnv()
	if err != nil {
		panic("Failed to load .env file")
	}
}