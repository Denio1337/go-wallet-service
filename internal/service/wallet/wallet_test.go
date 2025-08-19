package wallet_test

import (
	"testing"

	"github.com/Denio1337/go-wallet-service/internal/service/wallet"
	sv "github.com/Denio1337/go-wallet-service/internal/service/wallet"
	st "github.com/Denio1337/go-wallet-service/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) UpdateWallet(id uint, amount int) (uint, error) {
	args := m.Called(id, amount)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockStorage) GetWalletBalance(id uint) (uint, error) {
	args := m.Called(id)
	return args.Get(0).(uint), args.Error(1)
}

func TestWalletService_Deposit(t *testing.T) {
	storage := new(MockStorage)
	service := wallet.New(storage)
	id := uint(1)
	amount := 100

	storage.On("UpdateWallet", id, amount).Return(uint(amount), nil)

	params := &wallet.UpdateParams{
		ID:            id,
		Amount:        uint(amount),
		OperationType: wallet.DepositOperation,
	}

	wallet, err := service.Update(params)

	assert.NotNil(t, wallet)
	assert.NoError(t, err)
	assert.Equal(t, uint(amount), wallet.Amount)
	storage.AssertExpectations(t)
}

func TestWalletService_WithdrawSuccess(t *testing.T) {
	storage := new(MockStorage)
	service := wallet.New(storage)
	id := uint(1)
	amount := 100

	storage.On("UpdateWallet", id, -amount).Return(uint(0), nil)

	params := &wallet.UpdateParams{
		ID:            id,
		Amount:        uint(amount),
		OperationType: wallet.WithdrawOperation,
	}

	wallet, err := service.Update(params)

	assert.NotNil(t, wallet)
	assert.NoError(t, err)
	assert.Equal(t, uint(0), wallet.Amount)
	storage.AssertExpectations(t)
}

func TestWalletService_WithdrawError(t *testing.T) {
	storage := new(MockStorage)
	service := sv.New(storage)
	id := uint(1)
	amount := 100

	storage.On("UpdateWallet", id, -amount).Return(uint(0), st.ErrInvalidOperation)

	params := &sv.UpdateParams{
		ID:            id,
		Amount:        uint(amount),
		OperationType: sv.WithdrawOperation,
	}

	wallet, err := service.Update(params)

	assert.Nil(t, wallet)
	assert.Error(t, err)
	assert.Equal(t, sv.ErrInsufficientFunds, err)
	storage.AssertExpectations(t)
}
