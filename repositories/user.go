package repositories

import (
	"github.com/fajrizulfikar/ecommerce-api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) Create(user *models.User) (*models.User, error) {
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetById(userId uint) (*models.User, error) {
	var user models.User
	err := repo.DB.First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
