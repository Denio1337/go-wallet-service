package model

type Wallet struct {
	ID     uint `gorm:"primarykey"`
	Amount uint `gorm:"not null" json:"amount"`
}
