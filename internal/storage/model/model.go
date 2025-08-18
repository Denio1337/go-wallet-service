package model

import "gorm.io/gorm"

type (
	Wallet struct {
		gorm.Model
		Amount uint `gorm:"not null;index" json:"amount"`
	}
)
