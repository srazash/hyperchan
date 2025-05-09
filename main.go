package main

import (
	"fmt"
	"hyperchan/controllers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file\n")
	}
	fmt.Printf("%s\n", os.Getenv("GREETING"))

	controllers.Serve()
}
