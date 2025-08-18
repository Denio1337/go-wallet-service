package wallet

import (
	"github.com/Denio1337/go-wallet-service/internal/router/types/response"
	"github.com/Denio1337/go-wallet-service/internal/service/wallet"
	"github.com/gofiber/fiber/v2"
)

type UpdateDTO struct {
	WalletID      uint   `json:"walletID" validate:"required,gt=0"`
	OperationType string `json:"operationType" validate:"required,oneof=WITHDRAW DEPOSIT"`
	Amount        uint   `json:"amount" validate:"required,gt=0"`
}

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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(response.SuccessResponse(result))
}

// Update wallet
func Update(c *fiber.Ctx) error {
	// Parse body
	dto := new(UpdateDTO)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "incorrect input")
	}

	// Route to service
	result, err := wallet.Update(&wallet.UpdateParams{
		ID:            dto.WalletID,
		Amount:        dto.Amount,
		OperationType: wallet.OperationType(dto.OperationType),
	})

	// Handle service error
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(response.SuccessResponse(result))
}
