package db

import (
	"fmt"
	"gym/pkg/models"
	"gym/pkg/utils"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	psqlinfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", os.Getenv("DBHost"), os.Getenv("DBUser"), os.Getenv("DBName"), os.Getenv("DBPort"), os.Getenv("DBPassword"))
	db, err := gorm.Open(postgres.Open(psqlinfo), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB : %w", err)
	}
	db.AutoMigrate(&models.User{}, &models.Membership{}, &models.Admins{})
	if err := initializeSuperAdmin(db); err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func initializeSuperAdmin(db *gorm.DB) error {
	var count int64
	db.Model(&models.Admins{}).Where("role = ?", "superadmin").Count(&count)
	if count == 0 {
		superadmin := models.Admins{
			Email:    "superadmin@gmail.com",
			Password: utils.HashPassword("superadmin_password"),
			Role:     "superadmin",
		}
		if err := db.Create(&superadmin).Error; err != nil {
			return fmt.Errorf("failed to create superadmin user: %v", err)
		}
		fmt.Println("Superadmin user created successfully")
	} else {
		fmt.Println("Superadmin user already exists")
	}
	return nil
}
