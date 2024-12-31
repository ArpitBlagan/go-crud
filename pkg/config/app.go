package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (db * gorm.DB)

// func Connect() {
// 	var err error
// 	// Update connection string as per your database credentials
// 	//"postgres://default:6lJL8wOeSpIB@ep-spring-sun-a4w3r2o1.us-east-1.aws.neon.tech:5432/verceldb?sslmode=require"
// 	dsn := "postgres://default:6lJL8wOeSpIB@ep-spring-sun-a4w3r2o1.us-east-1.aws.neon.tech:5432/verceldb?sslmode=require	"
// 	db, err = gorm.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	fmt.Println("Database connected successfully!")
// }	
func Connect() {
	var err error
	// Update connection string as per your database credentials
	dsn := "host=ep-spring-sun-a4w3r2o1.us-east-1.aws.neon.tech user=default password=6lJL8wOeSpIB dbname=verceldb port=5432 sslmode=verify-full"
	
	// Open a connection to the database
	db, err := gorm.Open("postgres",dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Database connected successfully!",db)
}

func GetDB() *gorm.DB{
	return db
}