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
	User    models.User         `json:"user"`
}

func TestRegisterUser(t *testing.T) {
	t.Run("Should return 201 and confirmation for valid input", func(t *testing.T) {
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
		assert.Equal(t, 0, len(result.Errors))
	})

	t.Run("User saved in database", func(t *testing.T) {
		newUser := models.SignupInput{
			Username: "janedoe",
			Email:    "jane@doe.com",
			Password: "password",
		}

		writer := makeRequest("POST", "/register", newUser, false)

		var result ResponseModel
		json.Unmarshal(writer.Body.Bytes(), &result)

		assert.Equal(t, http.StatusCreated, writer.Code)

		assert.NotZero(t, result.User.ID, "User ID should not be 0")
		assert.False(t, result.User.CreatedAt.IsZero(), "CreatedAt should not be the zero timestamp")
	})

	t.Run("Verification email is sent to the user", func(t *testing.T) {
		newUser := models.SignupInput{
			Username: "johndoe",
			Email:    "john@doe.com",
			Password: "password",
		}

		writer := makeRequest("POST", "/register", newUser, false)

		var result ResponseModel
		json.Unmarshal(writer.Body.Bytes(), &result)

		assert.Equal(t, http.StatusCreated, writer.Code)
		assert.Equal(t, "We sent an email with a verification code to "+newUser.Email, result.Message)
	})
}
