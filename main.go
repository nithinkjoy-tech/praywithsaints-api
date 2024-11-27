package main

import (
	"fmt"
	"log"
	"prayer-book/config"
	"prayer-book/routes"
)

func main() {
	fmt.Println("Praise God!")
	config.ConnectDatabase()
	if err := config.InitializePrayersTable(); err != nil {
		log.Fatal(err)
	}

	defer config.DB.Close()

	router := routes.SetupRouter()
	router.Run(":8080")
}
