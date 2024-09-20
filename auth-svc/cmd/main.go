package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
