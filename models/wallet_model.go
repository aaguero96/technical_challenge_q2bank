package models

type WalletModel struct {
	WalletID int     `gorm:"column:wallet_id;primaryKey;autoIncrement:true"`
	Amount   float64 `gorm:"column:amount"`
}

func (WalletModel) TableName() string {
	return "wallets"
}
