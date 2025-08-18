package postgres

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Denio1337/go-wallet-service/internal/config"
	"github.com/Denio1337/go-wallet-service/internal/storage/contract"
	"github.com/Denio1337/go-wallet-service/internal/storage/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresStorage struct {
	db *gorm.DB
}

// Create new PostgreSQL storage implementation
func New() (contract.Storage, error) {
	// Parse port from environment
	p := config.Get(config.EnvDBPort)
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return nil, err
	}

	// Define data source
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Get(config.EnvDBHost),
		port,
		config.Get(config.EnvDBUser),
		config.Get(config.EnvDBPassword),
		config.Get(config.EnvDBName),
	)

	// Try to connect with default gorm config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate schemas to database
	db.AutoMigrate(&model.Wallet{})

	// Configure GORM connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute)

	return &PostgresStorage{db: db}, nil
}

// func (s *PostgresStorage) UpdateWallet(wallet *model.Wallet) (*model.Wallet, error) {
// 	// Create or update wallet
// 	err := s.db.Save(wallet).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return wallet, nil
// }

func (s *PostgresStorage) UpdateWallet(wallet *model.Wallet) (*model.Wallet, error) {
	var updated *model.Wallet

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var dbWallet model.Wallet
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&dbWallet, wallet.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				dbWallet = *wallet
			} else {
				return err
			}
		} else {
			dbWallet.Amount = wallet.Amount
		}

		if err := tx.Save(&dbWallet).Error; err != nil {
			return err
		}
		updated = &dbWallet
		return nil
	})

	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *PostgresStorage) GetWalletByID(id uint) (*model.Wallet, error) {
	var wallet model.Wallet

	err := s.db.First(&wallet, id).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, contract.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
