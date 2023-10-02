package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fajrizulfikar/ecommerce-api/database"
	"github.com/fajrizulfikar/ecommerce-api/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func setup() {
	errorEnv := godotenv.Load("../.env.test.local")
	if errorEnv != nil {
		panic("Failed to load env file")
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

	r.HandleFunc("/register", RegisterUser).Methods("POST")

	return r
}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}
