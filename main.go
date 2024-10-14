package main

import (
	"log"
	"muslim_pro/config"
	"muslim_pro/routes"
	"muslim_pro/seeders"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	// seeders.Seed()
	seeders.SeederCaption()

	r := routes.SetupRouter()
	r.Static("/public", "./public")
	r.Run(":8080")
}
