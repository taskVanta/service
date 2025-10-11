package main

import (
	"log"
	"service/config"
	handlers "service/handlers/users"
	models "service/models/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)

	r.Run(":8080")
}
