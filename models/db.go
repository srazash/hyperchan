package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBVersion() string {
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database\n")
	}

	var v string
	db.Raw("SELECT version()").Scan(&v)

	return fmt.Sprintf("PostgreSQL Version: %s\n", v)
}
