package main

import (
	"log"

	"github.com/Marlon-Sbardelatti/payment-service/internal/domain/payment"
	"github.com/Marlon-Sbardelatti/payment-service/internal/infrastructure/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db, err := database.NewConnection()
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	db.AutoMigrate(&payment.Payment{})
}
