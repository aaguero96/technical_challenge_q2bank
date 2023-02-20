package registerTypeRepository

import (
	"github.com/aaguero96/technical_challenge_q2bank/models"
	"gorm.io/gorm"
)

type registerTypeRepository struct {
	db *gorm.DB
}

func NewRegisterTypeRepository(db *gorm.DB) registerTypeRepository {
	return registerTypeRepository{
		db: db,
	}
}

func (rtr registerTypeRepository) GetAll() ([]models.RegisterTypeModel, error) {
	var registerTypes []models.RegisterTypeModel
	result := rtr.db.Find(&registerTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return registerTypes, nil
}
