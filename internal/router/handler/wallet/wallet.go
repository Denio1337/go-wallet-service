package wallet

import (
	"github.com/Denio1337/go-wallet-service/internal/router/types/response"
	"github.com/Denio1337/go-wallet-service/internal/service/wallet"
	"github.com/gofiber/fiber/v2"
)

// Get wallet by ID
func GetByID(c *fiber.Ctx) error {
	// Parse wallet ID from URL
	id, _ := c.ParamsInt("id", 0)
	if id <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "incorrect wallet ID specified")
	}

	// Route to service
	result, err := wallet.GetByID(&wallet.GetByIDParams{
		ID: uint(id),
	})

	// Handle service error
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "service error")
	}

	return c.JSON(response.SuccessResponse(result))
}

// Update wallet
func Update(c *fiber.Ctx) error {
	return c.JSON(response.SuccessResponse("test"))
}
