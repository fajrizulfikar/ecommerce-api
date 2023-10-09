package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fajrizulfikar/ecommerce-api/database"
	"github.com/fajrizulfikar/ecommerce-api/initializers"
	"github.com/fajrizulfikar/ecommerce-api/models"
	"github.com/fajrizulfikar/ecommerce-api/repositories"
	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func setup() {
	_, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	database.ConnectDB()
	database.Database.AutoMigrate(&models.User{})
}

func teardown() {
	migrator := database.Database.Migrator()
	migrator.DropTable(&models.User{})
}

func router() http.Handler {
	r := mux.NewRouter()

	userRepo := repositories.NewUserRepository(database.Database)

	authController := NewAuthController(userRepo)

	r.HandleFunc("/register", authController.RegisterUser).Methods("POST")

	return r
}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}
