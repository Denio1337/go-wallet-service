package wallet

import (
	"errors"

	"github.com/Denio1337/go-wallet-service/internal/storage"
)

type (
	GetByIDParams struct {
		ID uint
	}

	GetByIDResult struct {
		ID     uint
		Amount uint
	}

	UpdateParams struct {
		ID            uint
		Amount        uint
		OperationType OperationType
	}

	UpdateResult struct {
		ID     uint
		Amount uint
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

func GetByID(params *GetByIDParams) (*GetByIDResult, error) {
	wallet, err := storage.GetWalletByID(params.ID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &GetByIDResult{
		ID:     wallet.ID,
		Amount: wallet.Amount,
	}, nil
}

func Update(params *UpdateParams) (*UpdateResult, error) {
	// Negative amount if operation is withdraw
	amount := int(params.Amount)
	if params.OperationType == WithdrawOperation {
		amount = -amount
	}

	// Update wallet in DB
	newAmount, err := storage.UpdateWallet(params.ID, amount)
	if err != nil {
		if errors.Is(err, storage.ErrInvalidOperation) {
			return nil, ErrInsufficientFunds
		}

		return nil, err
	}

	return &UpdateResult{
		ID:     params.ID,
		Amount: newAmount,
	}, nil
}
