package wallet

import (
	"github.com/Denio1337/go-wallet-service/internal/router/types/response"
	"github.com/gofiber/fiber/v2"
)

// Get wallet by ID
func GetByID(c *fiber.Ctx) error {
	return c.JSON(response.SuccessResponse("test"))
}

// Update wallet
func Update(c *fiber.Ctx) error {
	return c.JSON(response.SuccessResponse("test"))
}
