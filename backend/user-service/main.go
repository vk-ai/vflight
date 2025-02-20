package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vk-ai/vflight/commons/database"
	"github.com/vk-ai/vflight/commons/logger"
	"github.com/vk-ai/vflight/user-service/config"
	"github.com/vk-ai/vflight/user-service/repository"
	"github.com/vk-ai/vflight/user-service/routes"
	"github.com/vk-ai/vflight/user-service/service"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger := logger.NewLogger(cfg.Environment)

	// Initialize database
	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database", err)
	}

	// Initialize repositories and services
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, logger)

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, userService, logger)

	// Start server
	logger.Info("Starting User service", map[string]interface{}{
		"port": cfg.Server.Port,
	})

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		logger.Fatal("Failed to start server", err)
	}
}
