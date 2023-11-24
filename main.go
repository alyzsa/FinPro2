package main

import (
	"os"

	"github.com/alyzsa/FinPro2/database"
)

func main() {
	database.StartDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
}
