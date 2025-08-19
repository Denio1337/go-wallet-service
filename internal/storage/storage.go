package storage

import (
	"github.com/Denio1337/go-wallet-service/internal/config"
	"github.com/Denio1337/go-wallet-service/internal/storage/contract"
	"github.com/Denio1337/go-wallet-service/internal/storage/postgres"
)

// Errors
var (
	ErrNotFound         = contract.ErrNotFound
	ErrInvalidOperation = contract.ErrInvalidOperation
)

func New(cfg *config.StorageConfig) contract.Storage {
	return postgres.New(cfg)
}
