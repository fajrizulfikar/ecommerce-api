package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fajrizulfikar/ecommerce-api/controllers"
	"github.com/fajrizulfikar/ecommerce-api/database"
	"github.com/fajrizulfikar/ecommerce-api/models"
	"github.com/fajrizulfikar/ecommerce-api/repositories"
	"github.com/fajrizulfikar/ecommerce-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
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

func serveApplication() {
	fmt.Println("Server running on port 3000")

	userRepo := repositories.NewUserRepository(database.Database)

	authController := controllers.NewAuthController(userRepo)

	log.Fatal(http.ListenAndServe(":3000", routes.Routes(authController)))
}
