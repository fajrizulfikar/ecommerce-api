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
}

func TestRegister(t *testing.T) {
	newUser := models.AuthenticationInput{
		Username: "johndoe",
		Email:    "john@doe.com",
		Password: "password",
	}

	writer := makeRequest("POST", "/register", newUser, false)

	var result ResponseModel
	json.Unmarshal(writer.Body.Bytes(), &result)

	assert.Equal(t, http.StatusCreated, writer.Code)
	assert.Equal(t, "User created!", result.Message)
	assert.Equal(t, 0, len(result.Errors))
}
