package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vk-ai/vflight/commons/database"
	"github.com/vk-ai/vflight/commons/logger"
	"go.uber.org/zap"
)

// SetupRouter initializes all the routes for the user service
func SetupRouter(db *database.Database) *gin.Engine {
	router := gin.Default()

	// Health check route
	router.GET("/health", healthCheck)

	// API routes group
	v1 := router.Group("/api/v1")
	{
		// Users routes group
		users := v1.Group("/users")
		{
			// Add your user routes here
			// users.POST("/", handlers.CreateUser)
			// users.GET("/:id", handlers.GetUser)
			// users.PUT("/:id", handlers.UpdateUser)
			// users.DELETE("/:id", handlers.DeleteUser)
		}
		logger.Info("Routes initialized successfully", zap.String("routes", users.BasePath()))
	}

	logger.Info("Routes initialized successfully")
	return router
}

// healthCheck handler for the health check endpoint
func healthCheck(c *gin.Context) {
	logger.Info("Health check requested")
	c.JSON(200, gin.H{
		"status":  "healthy",
		"service": "user-service",
	})
}
