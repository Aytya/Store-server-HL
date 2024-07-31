package main

import (
	"e-commerce"
	"e-commerce/internal/handler"
	"e-commerce/internal/repository"
	"log"
	"os"
)

func main() {
	dbURL := "postgres://postgres:7212Hey)@localhost:db/store"
	db := e_commerce.ConnectToDatabase(dbURL)
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error getting raw database object: %v\n", err)
		}
		sqlDB.Close()
	}()

	e_commerce.AutoMigrate(db)

	order := repository.NewOrderRepository(db)
	user := repository.NewUserRepository(db)
	product := repository.NewProductRepository(db)
	payment := repository.NewPaymentRepository(db)
	handlers := handler.NewHandler(order, payment, user, product)

	router := handlers.InitRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
