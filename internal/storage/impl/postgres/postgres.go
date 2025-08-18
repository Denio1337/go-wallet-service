package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Denio1337/go-wallet-service/internal/config"
	"github.com/Denio1337/go-wallet-service/internal/storage/contract"
	"github.com/Denio1337/go-wallet-service/internal/storage/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (s *PostgresStorage) UpdateWallet(id uint, amount int) (uint, error) {
	var newAmount uint

	err := s.db.Transaction(func(tx *gorm.DB) error {
		row := tx.Raw(`
      INSERT INTO wallets (id, amount) 
			VALUES (?, ?)
      ON CONFLICT (id) DO UPDATE
      	SET amount = wallets.amount + EXCLUDED.amount
      	WHERE wallets.amount + EXCLUDED.amount >= 0
      RETURNING amount;
  `, id, amount).Row()

		if scanErr := row.Scan(&newAmount); scanErr != nil {
			return contract.ErrInvalidOperation
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return newAmount, nil
}

func (s *PostgresStorage) GetWalletBalance(id uint) (uint, error) {
	var balance uint

	row := s.db.Raw(`SELECT amount FROM wallets WHERE id = ?;`, id).Row()

	if err := row.Scan(&balance); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, contract.ErrNotFound
		}

		return 0, err
	}
	return balance, nil
}
