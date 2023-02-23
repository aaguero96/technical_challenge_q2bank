package user_type_repository

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

func (utr userTypeRepository) GetById(id int) (models.UserTypeModel, error) {
	var userType models.UserTypeModel
	result := utr.db.First(&userType, id)
	if result.Error != nil {
		return models.UserTypeModel{}, result.Error
	}
	return userType, nil
}
