package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fajrizulfikar/ecommerce-api/models"
	"github.com/fajrizulfikar/ecommerce-api/repositories"
	"github.com/fajrizulfikar/ecommerce-api/utils"
)

type AuthController struct {
	Repo *repositories.UserRepository
}

func NewAuthController(repo *repositories.UserRepository) *AuthController {
	return &AuthController{Repo: repo}
}

func (ctrl *AuthController) RegisterUser(w http.ResponseWriter, req *http.Request) {
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

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   "Internal Server Error",
			"message": "Failed to process the request. Please try again later.",
		})
		return
	}

	newUser := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	savedUser, err := ctrl.Repo.Create(&newUser)
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
