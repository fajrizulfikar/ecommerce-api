package repositories

import (
	"github.com/fajrizulfikar/ecommerce-api/src/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(username string, password string, email string) error
}

type Repository struct {
	db *gorm.DB
}

func (repo Repository) RegisterUser(user models.User) error {
	result := repo.db.Create(&user)

	return result.Error
}
