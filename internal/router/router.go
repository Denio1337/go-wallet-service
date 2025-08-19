package router

import (
	"errors"

	"github.com/Denio1337/go-wallet-service/internal/config"
	"github.com/Denio1337/go-wallet-service/internal/router/handler/wallet"
	"github.com/Denio1337/go-wallet-service/internal/router/types"
	"github.com/Denio1337/go-wallet-service/internal/router/types/response"
	sw "github.com/Denio1337/go-wallet-service/internal/service/wallet"
	"github.com/Denio1337/go-wallet-service/internal/storage/contract"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	App     *fiber.App
	Storage contract.Storage
	Config  *config.Config
}

func (r *Router) Serve() error {
	return r.App.Listen(r.Config.Address)
}

// Create and configure application router
func New(cfg *config.Config, storage contract.Storage) *Router {
	// Initialize Fiber application
	fiberApp := fiber.New(fiber.Config{
		Prefork:      true,                // Spawn multiple Go processes listening on the same port
		ServerHeader: "Go Wallet Service", // Set "Server" HTTP-header
		AppName:      "Go Wallet Service",
		ErrorHandler: handleError,
	})

	// Inject storage and config into the router
	router := &Router{
		App:     fiberApp,
		Storage: storage,
		Config:  cfg,
	}

	// Configure endpoints
	setupRoutes(router)

	return router
}

// Set router api
func setupRoutes(router *Router) {
	// General API group
	apiGroup := router.App.Group(types.ApiPath, logger.New())

	// Wallet group
	walletGroup := apiGroup.Group(types.WalletsPath)
	walletService := sw.New(router.Storage)
	walletHandler := wallet.New(walletService)
	walletGroup.Get("/:"+types.WalletIDParam, walletHandler.GetByID)
	walletGroup.Post("/", walletHandler.Update)
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
