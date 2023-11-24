package main

import (
	"github.com/alyzsa/FinPro2/database"
	"github.com/alyzsa/FinPro2/router"
	"os"
)

func main() {
	database.StartDB()

	// Get the port from the environment variable or default to "8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := router.StartApp()
	r.Run(":" + port)
}
