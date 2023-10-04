package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fajrizulfikar/ecommerce-api/models"
)

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	var input models.SignupInput
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	hasErrors := false
	var errors []map[string]string

	if input.Username == "" {
		errors = append(errors, map[string]string{"username": "Username not received"})
		hasErrors = true
	}
	if input.Email == "" {
		errors = append(errors, map[string]string{"email": "Email not received"})
		hasErrors = true
	}
	if input.Password == "" {
		errors = append(errors, map[string]string{"password": "Password not received"})
		hasErrors = true
	}

	if hasErrors {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Invalid input",
			"errors":  errors,
		})
		return
	}

	newUser := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	savedUser, err := newUser.Create()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Failed create user",
			"errors":  errors,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created!",
		"user":    savedUser,
	})
}
