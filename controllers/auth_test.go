package controllers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/fajrizulfikar/ecommerce-api/models"
	"github.com/stretchr/testify/assert"
)

type ResponseModel struct {
	Message string              `json:"message"`
	Errors  []map[string]string `json:"errors"`
	User    models.User
}

func TestCreateUserReturnsCreatedStatus(t *testing.T) {
	newUser := models.SignupInput{
		Username: "johndoe",
		Email:    "john@doe.com",
		Password: "password",
	}

	writer := makeRequest("POST", "/register", newUser, false)

	var result ResponseModel
	json.Unmarshal(writer.Body.Bytes(), &result)

	assert.Equal(t, http.StatusCreated, writer.Code)
	assert.Equal(t, "User created!", result.Message)

	assert.NotZero(t, result.User.ID, "User ID should not be 0")
	assert.False(t, result.User.CreatedAt.IsZero(), "CreatedAt should not be the zero timestamp")
}
