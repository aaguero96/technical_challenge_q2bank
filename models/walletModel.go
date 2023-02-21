package models

type WalletModel struct {
	WalletID int     `gorm:"column:wallet_id;primaryKey"`
	Amount   float64 `gorm:"column:amount"`
}

func (WalletModel) TableName() string {
	return "wallets"
}
