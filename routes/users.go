// service/routes/users.go
package routes

import (
	doc "service/handlers/docs"
	userHandler "service/handlers/users"
	"service/middleware"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // replace with your frontend origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply CORS middleware globally
	router.Use(cors.New(config))
	router.GET("/", doc.ServeAPIDocs)
	router.GET("/docs", doc.ServeAPIDocs)
	userGroup := router.Group("/api/users")

	userGroup.POST("/signin", userHandler.Signin)
	userGroup.POST("/signup", userHandler.Signup)
	userGroup.GET("/logout", userHandler.Logout)

	protected := userGroup.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", userHandler.ProfileHandler)
		protected.POST("/", userHandler.CreateUser)
		protected.GET("/", userHandler.GetUsers)
		protected.GET("/:id", userHandler.GetUserByID)
		protected.PUT("/:id", userHandler.UpdateUser)
		protected.DELETE("/:id", userHandler.DeleteUser)
	}

}
