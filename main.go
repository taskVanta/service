package main

import (
	"service/config"
	models "service/models/users"
	"service/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	routes.RegisterUserRoutes(r)

	r.Run(":8080")
}
