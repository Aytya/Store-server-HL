package e_commerce

import (
	"e-commerce/internal/domain"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDatabase(dbURL string) *gorm.DB {
	var err error
	db, err = gorm.Open(postgres.Open(dbURL+"?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting database connection: %v\n", err)
	}

	for retries := 1; retries <= 5; retries++ {
		err = sqlDB.Ping()
		if err != nil {
			log.Printf("Error pinging database (retry %d): %v\n", retries, err)
			time.Sleep(5 * time.Second)
		} else {
			log.Println("Successfully connected to the database")
			break
		}
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Product{}, &domain.User{}, &domain.Order{}, &domain.Payment{})
	if err != nil {
		log.Fatalf("Error during migration: %v\n", err)
	}
}
