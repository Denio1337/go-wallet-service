package contract

import (
	"errors"

	"github.com/Denio1337/go-wallet-service/internal/storage/model"
)

type Storage interface {
	UpdateWallet(*model.Wallet) (*model.Wallet, error)
	GetWalletByID(uint) (*model.Wallet, error)
}

var (
	ErrNotFound = errors.New("wallet not found")
)
