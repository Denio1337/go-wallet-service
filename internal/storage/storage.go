package storage

import (
	"github.com/Denio1337/go-wallet-service/internal/storage/contract"
	"github.com/Denio1337/go-wallet-service/internal/storage/impl"
)

// Global storage instance
var instance contract.Storage

// Create DB Connection with current implementation
func init() {
	instance = impl.Impl
}

// Interface

func GetWalletBalance(id uint) (uint, error) {
	return instance.GetWalletBalance(id)
}

func UpdateWallet(id uint, amount int) (uint, error) {
	return instance.UpdateWallet(id, amount)
}

// Errors

var (
	ErrNotFound         = contract.ErrNotFound
	ErrInvalidOperation = contract.ErrInvalidOperation
)
