package initializers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvVariables() {
	// Check if running in production environment
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}
}
