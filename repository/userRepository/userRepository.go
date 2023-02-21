package userRepository

import (
	"errors"

	"github.com/aaguero96/technical_challenge_q2bank/models"
	"github.com/aaguero96/technical_challenge_q2bank/utils"
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

func (ur userRepository) CreateUser(name, email, password string, registerNumber int64, registerTypeID, userTypeID int) (models.UserModel, error) {
	// Create transaction
	tx := ur.db.Begin()

	// Create new wallet with amount zero
	var newWallet models.WalletModel
	newWallet.Amount = 0
	result := tx.Model(&models.WalletModel{}).Create(&newWallet)
	if result.Error != nil {
		return models.UserModel{}, result.Error
	}

	// Verify if registerNumber has correct number of chars
	var registerType models.RegisterTypeModel
	result = tx.Model(&models.RegisterTypeModel{}).First(&registerType, registerTypeID)
	if result.Error != nil {
		return models.UserModel{}, result.Error
	}
	if err := utils.ValidateRegisterNumber(registerNumber, registerType.RegisterType); err != nil {
		return models.UserModel{}, errors.New("register number is invalid, that was considering the register type passed")
	}

	// hash password
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return models.UserModel{}, errors.New("internal error, related to password input")
	}

	// Create user
	newUser := models.UserModel{
		Name:           name,
		RegisterNumber: registerNumber,
		RegisterTypeID: registerTypeID,
		Email:          email,
		WalletID:       newWallet.WalletID,
		UserTypeID:     userTypeID,
		Password:       hashPassword,
	}
	result = tx.Model(&models.UserModel{}).Create(&newUser)
	if result.Error != nil {
		return models.UserModel{}, result.Error
	}

	// End Transaction
	tx.Commit()

	return newUser, nil
}
