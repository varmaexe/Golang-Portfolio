package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file %v", err)
	}
}
