package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fajrizulfikar/ecommerce-api/controllers"
	"github.com/fajrizulfikar/ecommerce-api/database"
	"github.com/fajrizulfikar/ecommerce-api/initializers"
	"github.com/fajrizulfikar/ecommerce-api/models"
	"github.com/fajrizulfikar/ecommerce-api/repositories"
	"github.com/fajrizulfikar/ecommerce-api/routes"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	_, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
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
