package db

import (
	"fmt"
	"github.com/karthik11135/golang-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDb() bool {
	dsn := "postgresql://postgres:mysecretpassword@localhost:5432/postgres"
	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return false
	}
	fmt.Println("Successfully connected to the db")
	db = pg

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Playlist{})

	return true
}

func GetDb() *gorm.DB {
	return db
}