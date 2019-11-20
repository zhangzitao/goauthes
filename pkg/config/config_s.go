package config

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadConfig is for configuring
func LoadConfig() bool {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
	return true
}
