package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name           string `gorm:"column:name"`
	RegisterNumber int64  `gorm:"column:register_number;index:register,unique"`
	RegisterTypeID int    `gorm:"column:register_type_id;index:register,unique"`
	Email          string `gorm:"column:email;unique"`
	WalletID       int    `gorm:"column:wallet_id"`
	UserTypeID     int    `gorm:"column:user_type_id"`
	Password       string `gorm:"password"`
}

func (UserModel) TableName() string {
	return "users"
}
