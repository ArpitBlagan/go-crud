package config

import {
	"github.com/jinzhu/gorm",
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"fmt"
}

var (db * gorm.DB)

func Connect() {
	var err error
	// Update connection string as per your database credentials
	dsn := "host=localhost user=your_user password=your_password dbname=your_dbname port=5432 sslmode=disable"
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Database connected successfully!")
}	

func GetDB() *gorm.DB{
	return db
}