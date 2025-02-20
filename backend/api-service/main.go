package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vk-ai/vflight/api-service/config"
	"github.com/vk-ai/vflight/api-service/routes"
	"github.com/vk-ai/vflight/commons/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger := logger.NewLogger(cfg.Environment)

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, cfg, logger)

	// Start server
	logger.Info("Starting API service", map[string]interface{}{
		"port": cfg.Server.Port,
	})

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		logger.Fatal("Failed to start server", err)
	}
}
