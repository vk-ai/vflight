package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vk-ai/vflight/commons/database"
	"github.com/vk-ai/vflight/commons/logger"
)

type Handler struct {
	db *database.Database
}

func NewHandler(db *database.Database) *Handler {
	return &Handler{
		db: db,
	}
}

// HealthCheck handles the health check endpoint
func (h *Handler) HealthCheck(c *gin.Context) {
	logger.Info("Health check requested")
	c.JSON(200, gin.H{
		"status":  "healthy",
		"service": "api-service",
	})
}
