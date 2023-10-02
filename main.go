package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fajrizulfikar/ecommerce-api/database"
	"github.com/fajrizulfikar/ecommerce-api/models"
	"github.com/fajrizulfikar/ecommerce-api/routes"
)

func main() {
	loadDatabase()
	serverApplication()
}

func loadDatabase() {
	database.ConnectDB()
	database.Database.AutoMigrate(&models.User{})
}

func serverApplication() {
	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", routes.Routes()))
}
