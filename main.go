package main

import (
	"github.com/alyzsa/FinPro2/database"
	"github.com/alyzsa/FinPro2/router"
	
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
