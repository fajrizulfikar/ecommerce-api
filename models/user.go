package models

import (
	"time"

	"github.com/fajrizulfikar/ecommerce-api/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null;" json:"-"`
	Email     string    `gorm:"size:255;not null;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
}

func (user *User) Create() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
