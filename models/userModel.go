package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name           string `gorm:"column:name"`
	RegisterNumber int64  `gorm:"column:register_number"`
	RegisterTypeID int    `gorm:"column:register_type_id"`
	Email          string `gorm:"column:email"`
	WalletID       int    `gorm:"column:wallet_id"`
	UserTypeID     int    `gorm:"column:user_type_id"`
}

func (UserModel) TableName() string {
	return "users"
}
