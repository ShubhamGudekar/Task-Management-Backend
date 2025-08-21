package infrastructure

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load("internal/infrastructure/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
