package models

import (
	"gorm.io/gorm"
)

type TransactionModel struct {
	gorm.Model
	PayerID int     `gorm:"column:payer_id"`
	PayeeID int     `gorm:"column:payee_id"`
	Amount  float64 `gorm:"column:amount"`
	Status  string  `gorm:"column:status,type:status"`
}

func (TransactionModel) TableName() string {
	return "transactions"
}
