package wallet

import (
	"errors"

	"github.com/Denio1337/go-wallet-service/internal/router/types"
	cerror "github.com/Denio1337/go-wallet-service/internal/router/types/error"
	"github.com/Denio1337/go-wallet-service/internal/router/types/response"
	"github.com/Denio1337/go-wallet-service/internal/router/validator"
	"github.com/Denio1337/go-wallet-service/internal/service/wallet"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type handler struct {
	service wallet.Service
}

type (
	UpdateDTO struct {
		WalletID      uint   `json:"walletID" validate:"required,gt=0"`
		OperationType string `json:"operationType" validate:"required,oneof=WITHDRAW DEPOSIT"`
		Amount        uint   `json:"amount" validate:"required,gt=0"`
	}

	Wallet struct {
		ID     uint `json:"id"`
		Amount uint `json:"amount"`
	}
)

var (
	ErrWalletID = fiber.NewError(fiber.StatusBadRequest, "incorrect wallet ID specified")
	ErrUpdate   = fiber.NewError(fiber.StatusBadRequest, "incorrect input")
)

func New(service wallet.Service) Handler {
	return &handler{
		service: service,
	}
}

// Get wallet by ID
func (h *handler) GetByID(c *fiber.Ctx) error {
	// Parse wallet ID from URL
	id, _ := c.ParamsInt(types.WalletIDParam, types.WalletIDDefault)
	if id <= 0 {
		return ErrWalletID
	}

	// Route to service
	result, err := h.service.GetByID(uint(id))

	// Handle service error
	if err != nil {
		if errors.Is(err, wallet.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		return err
	}

	return c.JSON(response.SuccessResponse(&Wallet{
		ID:     result.ID,
		Amount: result.Amount,
	}))
}

// Update wallet
func (h *handler) Update(c *fiber.Ctx) error {
	// Parse body
	dto := new(UpdateDTO)
	if err := c.BodyParser(dto); err != nil {
		return ErrUpdate
	}

	// Validate body
	if errs := validator.Validate(dto); len(errs) > 0 {
		return cerror.ValidationError(errs)
	}

	// Route to service
	result, err := h.service.Update(&wallet.UpdateParams{
		ID:            dto.WalletID,
		Amount:        dto.Amount,
		OperationType: wallet.OperationType(dto.OperationType),
	})

	// Handle service error
	if err != nil {
		if errors.Is(err, wallet.ErrBadWithdraw) ||
			errors.Is(err, wallet.ErrInsufficientFunds) ||
			errors.Is(err, wallet.ErrInvalidOperation) {
			return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
		}

		return err
	}

	return c.JSON(response.SuccessResponse(&Wallet{
		ID:     result.ID,
		Amount: result.Amount,
	}))
}
