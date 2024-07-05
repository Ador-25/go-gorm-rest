package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func connect() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	var err error
	db, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")
	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Println("Initializing database connection...")
		db = connect()
	} else {
		log.Println("Reusing existing database connection")
	}
	return db
}
