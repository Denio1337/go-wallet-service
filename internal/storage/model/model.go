package model

import "gorm.io/gorm"

type (
	Wallet struct {
		gorm.Model
		Amount int `gorm:"not null;index" json:"amount"`
	}
)
