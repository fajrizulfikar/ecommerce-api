package main

import (
	"log"
	"net/http"

	"github.com/fajrizulfikar/ecommerce-api/config"
	"github.com/fajrizulfikar/ecommerce-api/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	log.Println("Server started at port 3000")
	log.Fatal(http.ListenAndServe(":3000", routes.Routes()))
}
