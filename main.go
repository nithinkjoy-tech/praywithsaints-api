package main

import (
	"fmt"
	"prayer-book/config"
	"prayer-book/routes"
)

func main() {
	fmt.Println("Praise God!")
	config.ConnectDatabase()
	config.InitializePrayersTable()

	defer config.DB.Close()

	router := routes.SetupRouter()
	router.Run(":8080")
}
