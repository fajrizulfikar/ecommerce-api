package repositories

import (
	"os"
	"testing"

	"github.com/fajrizulfikar/ecommerce-api/database"
	"github.com/fajrizulfikar/ecommerce-api/models"
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
