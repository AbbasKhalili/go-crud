package config

import (
	"fmt"
	"go-crud/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func setupDbConnection() {
	dsn := "sqlserver://localhost:1433?database=Cafe3&trusted_connection=yes"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	fmt.Println("Database connected")
}

func ConnectDatabase() {
	setupDbConnection()
	DB.AutoMigrate(&models.Product{})
}
