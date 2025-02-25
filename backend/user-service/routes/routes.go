package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vk-ai/vflight/commons/database"
	"github.com/vk-ai/vflight/commons/logger"
)

// SetupRouter initializes all the routes for the user service
func SetupRouter(db *database.Database) *gin.Engine {
	router := gin.Default()

	// Initialize handler
	handler := NewHandler(db)

	// Health check route
	router.GET("/health", handler.HealthCheck)

	// API routes group
	v1 := router.Group("/api/v1")
	{
		// Users routes group
		users := v1.Group("/users")
		{
			users.GET("/:id", handler.GetUser)
			users.POST("/", handler.CreateUser)
			users.PUT("/:id", handler.UpdateUser)
			users.DELETE("/:id", handler.DeleteUser)
		}
	}

	logger.Info("Routes initialized successfully")
	return router
}
