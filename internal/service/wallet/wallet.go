package wallet

import "github.com/Denio1337/go-wallet-service/internal/storage"

type (
	GetByIDParams struct {
		ID uint
	}

	GetByIDResult struct {
		ID     uint
		Amount uint
	}

	UpdateParams struct {
		ID     uint
		Amount uint
	}

	UpdateResult struct {
		ID     uint
		Amount uint
	}
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
	return nil, nil
}
