package main

import (
	"service/config"
	handlers "service/handlers/users"
	models "service/models/users"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)

	r.Run(":8080")
}
