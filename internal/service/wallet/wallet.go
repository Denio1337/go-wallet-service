package wallet

import (
	"errors"

	"github.com/Denio1337/go-wallet-service/internal/storage"
	"github.com/Denio1337/go-wallet-service/internal/storage/model"
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

func GetByID(params *GetByIDParams) (*GetByIDResult, error) {
	wallet, err := storage.GetWalletByID(params.ID)
	if err != nil {
		return nil, err
	}

	return &GetByIDResult{
		ID:     wallet.ID,
		Amount: wallet.Amount,
	}, nil
}

func Update(params *UpdateParams) (*UpdateResult, error) {
	// Get wallet with specified ID
	wallet, err := storage.GetWalletByID(params.ID)

	// Unexpected error
	if err != nil && err != storage.ErrNotFound {
		return nil, err
	}

	// Wallet was not found
	if err != nil {
		// Trying to withdraw unexisting wallet
		if params.OperationType == WithdrawOperation {
			return nil, errors.New("wallet not found for withdraw")
		}

		wallet = &model.Wallet{ID: params.ID, Amount: 0}
	}

	// Deposit or withdraw amount of wallet
	switch params.OperationType {
	case WithdrawOperation:
		if wallet.Amount < params.Amount {
			return nil, errors.New("insufficient funds")
		}
		wallet.Amount -= params.Amount
	case DepositOperation:
		wallet.Amount += params.Amount
	default:
		return nil, errors.New("invalid operation type")
	}

	// Update wallet in DB
	wallet, err = storage.UpdateWallet(wallet)
	if err != nil {
		return nil, err
	}

	return &UpdateResult{
		ID:     wallet.ID,
		Amount: wallet.Amount,
	}, nil
}
