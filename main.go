package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fajrizulfikar/ecommerce-api/database"
	"github.com/fajrizulfikar/ecommerce-api/models"
	"github.com/fajrizulfikar/ecommerce-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serverApplication()
}

func loadEnv() {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		panic("Failed to load env file")
	}
}

func loadDatabase() {
	database.ConnectDB()
	database.Database.AutoMigrate(&models.User{})
}

func serverApplication() {
	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", routes.Routes()))
}
