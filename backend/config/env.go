package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}