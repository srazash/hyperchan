package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBVersion() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file\n")
	}

	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database\n")
	}

	var v string
	db.Raw("SELECT version()").Scan(&v)

	return fmt.Sprintf("PostgreSQL Version: %s\n", v)
}
