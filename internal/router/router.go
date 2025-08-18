package router

import (
	"errors"

	"github.com/Denio1337/go-wallet-service/internal/router/handler/wallet"
	"github.com/Denio1337/go-wallet-service/internal/router/types/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Create and configure Fiber application
func New() *fiber.App {
	router := fiber.New(fiber.Config{
		Prefork:      true,                // Spawn multiple Go processes listening on the same port
		ServerHeader: "Go Wallet Service", // Set "Server" HTTP-header
		AppName:      "Go Wallet Service",
		ErrorHandler: handleError,
	})

	// Configure endpoints
	setupRoutes(router)

	return router
}

// Set router api
func setupRoutes(app *fiber.App) {
	// General API group
	apiGroup := app.Group("/api/v1", logger.New())

	// Wallet group
	walletGroup := apiGroup.Group("/wallets")
	walletGroup.Get("/:id", wallet.GetByID)
	walletGroup.Post("/", wallet.Update)
}

// Handle error response
func handleError(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(response.ErrorResponse(err.Error()))
}
