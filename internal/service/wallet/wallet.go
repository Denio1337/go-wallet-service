package wallet

import (
	"errors"

	"github.com/Denio1337/go-wallet-service/internal/storage"
	"github.com/Denio1337/go-wallet-service/internal/storage/contract"
)

type Service interface {
	GetByID(id uint) (*Wallet, error)
	Update(params *UpdateParams) (*Wallet, error)
}

type service struct {
	storage contract.Storage
}

type (
	Wallet struct {
		ID     uint
		Amount uint
	}

	UpdateParams struct {
		ID            uint
		Amount        uint
		OperationType OperationType
	}

	OperationType string
)

const (
	WithdrawOperation OperationType = "WITHDRAW"
	DepositOperation  OperationType = "DEPOSIT"
)

// Custom errors
var (
	ErrBadWithdraw       = errors.New("wallet not found for withdraw")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidOperation  = errors.New("invalid operation type")
	ErrNotFound          = errors.New("wallet not found")
)

func New(storage contract.Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) GetByID(id uint) (*Wallet, error) {
	amount, err := s.storage.GetWalletBalance(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &Wallet{
		ID:     id,
		Amount: amount,
	}, nil
}

func (s *service) Update(params *UpdateParams) (*Wallet, error) {
	// Negative amount if operation is withdraw
	amount := int(params.Amount)
	if params.OperationType == WithdrawOperation {
		amount = -amount
	}

	// Update wallet in DB
	newAmount, err := s.storage.UpdateWallet(params.ID, amount)
	if err != nil {
		if errors.Is(err, storage.ErrInvalidOperation) {
			return nil, ErrInsufficientFunds
		}

		return nil, err
	}

	return &Wallet{
		ID:     params.ID,
		Amount: newAmount,
	}, nil
}
