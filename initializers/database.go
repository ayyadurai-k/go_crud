package initializers

import (
	"fmt"
	"os"
	"log"
  	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	var err error

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_SSL_MODE := os.Getenv("DB_SSL_MODE")
	DB_TIMEZONE := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL_MODE, DB_TIMEZONE)
		
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	}
}