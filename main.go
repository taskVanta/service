package main

import (
	"log"
	"service/config"
	models "service/models/users"
	"service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	routes.RegisterUserRoutes(r)

	r.Run(":8080")
}
