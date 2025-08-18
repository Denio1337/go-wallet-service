package error

import (
	"fmt"
	"strings"

	"github.com/Denio1337/go-wallet-service/internal/router/validator"
	"github.com/gofiber/fiber/v2"
)

// Create fiber error about invalid validation
func ValidationError(errs []validator.ValidationError) *fiber.Error {
	errMsgs := make([]string, 0)

	// Formatting list of validation errors
	for _, err := range errs {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"[%s]: '%v' | Needs to implement '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}

	return &fiber.Error{
		Code:    fiber.ErrBadRequest.Code,
		Message: strings.Join(errMsgs, " and "),
	}
}
