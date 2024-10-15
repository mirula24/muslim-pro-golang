package main

import (
	"log"
	"muslim_pro/config"
	_ "muslim_pro/docs"
	"muslim_pro/routes"
	"muslim_pro/seeders"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
