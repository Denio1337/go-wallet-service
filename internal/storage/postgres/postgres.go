package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Denio1337/go-wallet-service/internal/config"
	"github.com/Denio1337/go-wallet-service/internal/storage/contract"
	"github.com/Denio1337/go-wallet-service/internal/storage/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresStorage struct {
	db *gorm.DB
}

// Create new PostgreSQL storage implementation
func New(cfg *config.StorageConfig) contract.Storage {
	// Define data source
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	// Try to connect with default gorm config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate schemas to database
	db.AutoMigrate(&model.Wallet{})

	// Configure GORM connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute)

	return &postgresStorage{db: db}
}

func (s *postgresStorage) UpdateWallet(id uint, amount int) (uint, error) {
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

func (s *postgresStorage) GetWalletBalance(id uint) (uint, error) {
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
