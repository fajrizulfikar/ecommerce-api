package repositories

import (
	"testing"

	"github.com/fajrizulfikar/ecommerce-api/database"
	"github.com/fajrizulfikar/ecommerce-api/models"
	"github.com/fajrizulfikar/ecommerce-api/utils"
	"github.com/stretchr/testify/assert"
)

func TestReturnHashedPasswordWhenUserCreated(t *testing.T) {
	newUser := models.User{
		Username: "johndoe",
		Email:    "john@doe.com",
		Password: "password",
	}

	userRepo := NewUserRepository(database.Database)

	savedUser, err := userRepo.Create(&newUser)
	if err != nil {
		t.Fatalf("Failed create user: %v", err)
	}

	fetchedUser, err := userRepo.GetById(savedUser.ID)
	if err != nil {
		t.Fatalf("Failed create user: %v", err)
	}

	isValidHash := utils.CheckPasswordHash(fetchedUser.Password, newUser.Password)
	assert.True(t, isValidHash, "Password was not hashed correctly")
}
