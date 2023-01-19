package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var connStr = "host=localhost user=postgres password=root dbname=gorm port=5432"

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connection successful")
	}
}
