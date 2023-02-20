package userRepository

import (
	"github.com/aaguero96/technical_challenge_q2bank/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{
		db: db,
	}
}

func (ur userRepository) GetAll() ([]models.UserModel, error) {
	var users []models.UserModel
	result := ur.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (ur userRepository) GetById(id int) (models.UserModel, error) {
	var user models.UserModel
	result := ur.db.First(&user, id)
	if result.Error != nil {
		return models.UserModel{}, result.Error
	}
	return user, nil
}
