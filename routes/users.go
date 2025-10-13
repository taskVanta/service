// service/routes/users.go
package routes

import (
	handlers "service/handlers/users"
	"service/middleware"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {

	config := cors.Config{
		AllowOrigins:     []string{"https://yourfrontend.com"}, // replace with your frontend origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply CORS middleware globally
	router.Use(cors.New(config))

	userGroup := router.Group("/api/users")

	userGroup.POST("/signin", handlers.Signin)
	userGroup.POST("/signup", handlers.Signup)

	protected := userGroup.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/", handlers.CreateUser)
		protected.GET("/", handlers.GetUsers)
		protected.GET("/:id", handlers.GetUserByID)
		protected.PUT("/:id", handlers.UpdateUser)
		protected.DELETE("/:id", handlers.DeleteUser)
	}
}
