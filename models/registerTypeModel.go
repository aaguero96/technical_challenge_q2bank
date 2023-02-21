package models

type RegisterTypeModel struct {
	RegisterTypeID int    `gorm:"column:register_type_id;primaryKey"`
	RegisterType   string `gorm:"column:register_type"`
}

func (RegisterTypeModel) TableName() string {
	return "register_types"
}
