package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDB() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		panic("Failed to parse port to number")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Assign Database here
	if err != nil {
		panic("Failed to connect to database")
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

func DisconnectDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	sqlDB.Close()
}
