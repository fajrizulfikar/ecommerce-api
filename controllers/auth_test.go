package controllers

import (
	"net/http"
	"testing"

	"github.com/fajrizulfikar/ecommerce-api/models"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	newUser := models.AuthenticationInput{
		Username: "john",
		Password: "password",
	}
	writer := makeRequest("POST", "/register", newUser, false)
	assert.Equal(t, http.StatusCreated, writer.Code)
}
