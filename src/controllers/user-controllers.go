package controllers

import (
	"net/http"

	"github.com/fajrizulfikar/ecommerce-api/src/models"
)

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	var user models.User

	user.Username = req.FormValue("username")
	user.Password = req.FormValue("password")
	user.Email = req.FormValue("email")
}
