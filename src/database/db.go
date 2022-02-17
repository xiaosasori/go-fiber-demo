package database

import (
	"go-fiber-demo/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := "host=db user=postgres password=123456 dbname=go-fiber-demo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{}, models.Product{})
}
