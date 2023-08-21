package db

import (
	"auth/internal/models" // Import your model package
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	dsn = fmt.Sprintf(dsn, os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	fmt.Println(dsn)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	err = MigrateModels(DB)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func MigrateModels(DB *gorm.DB) error {
	// AutoMigrate models
	err := DB.AutoMigrate(&models.User{},
		&models.Credientials{}, &models.ResourceGroup{},
		&models.UserPermissions{}, &models.ResourcePermission{})
	if err != nil {
		log.Fatal("Failed to auto migrate models:", err)
		return err
	}

	return nil
}
