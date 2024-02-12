package main

import (
	"auth-service/internal/app"
	"auth-service/internal/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("Starting auth-service's server")

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Load config from .env
	cfg := config.NewConfig()

	app.StartServer(cfg)
}
