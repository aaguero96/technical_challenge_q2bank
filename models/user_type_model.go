package models

type UserTypeModel struct {
	UserTypeID int    `gorm:"column:user_type_id;primaryKey"`
	UserType   string `gorm:"column:user_type"`
}

func (UserTypeModel) TableName() string {
	return "user_types"
}
