package models

type WalletModel struct {
	WalletID int `gorm:"column:wallet_id,primary_key"`
	Amount   int `gorm:"column:amount"`
}

func (WalletModel) TableName() string {
	return "wallets"
}
