package main

import (
	"log"

	"github.com/Denio1337/go-wallet-service/internal/config"
	"github.com/Denio1337/go-wallet-service/internal/router"
)

func main() {
	// Create application instance
	app := router.New()

	// Run application
	log.Fatal(app.Listen(config.Get(config.EnvAppAddress)))
}
