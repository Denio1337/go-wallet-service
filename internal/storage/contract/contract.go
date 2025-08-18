package contract

import (
	"errors"
)

type Storage interface {
	UpdateWallet(uint, int) (uint, error)
	GetWalletBalance(uint) (uint, error)
}

var (
	ErrNotFound         = errors.New("wallet not found")
	ErrInvalidOperation = errors.New("invalid wallet operation")
)
