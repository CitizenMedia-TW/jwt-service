package main

import (
	"github.com/joho/godotenv"
	"jwt-service/internal/app"
	"jwt-service/internal/config"
	"log"
)

func main() {
	log.Println("Starting jwt-service's server")

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Load config from .env
	cfg := config.NewConfig()

	app.StartServer(cfg)
}
