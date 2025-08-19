package main

import (
	"log"

	"github.com/Denio1337/go-wallet-service/internal/config"
	"github.com/Denio1337/go-wallet-service/internal/router"
	"github.com/Denio1337/go-wallet-service/internal/storage"
)

func main() {
	// Load configuration from environment variables
	cfg := config.MustLoad()

	// Initialize storage
	storage := storage.New(&cfg.StorageConfig)

	// Create application instance
	router := router.New(&cfg.AppConfig, storage)

	// Run application
	log.Fatal(router.Serve())
}
