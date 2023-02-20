package userTypeRepository

import (
	"github.com/aaguero96/technical_challenge_q2bank/models"
	"gorm.io/gorm"
)

type userTypeRepository struct {
	db *gorm.DB
}

func NewUserTypeRepository(db *gorm.DB) userTypeRepository {
	return userTypeRepository{
		db: db,
	}
}

func (utr userTypeRepository) GetAll() ([]models.UserTypeModel, error) {
	var userTypes []models.UserTypeModel
	result := utr.db.Find(&userTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return userTypes, nil
}
