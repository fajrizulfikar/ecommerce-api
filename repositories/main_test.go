package repositories

import (
	"log"
	"os"
	"testing"

	"github.com/fajrizulfikar/ecommerce-api/database"
	"github.com/fajrizulfikar/ecommerce-api/initializers"
	"github.com/fajrizulfikar/ecommerce-api/models"
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
