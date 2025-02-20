package main

import (
	cfg "github.com/vk-ai/vflight/commons/config"
	"github.com/vk-ai/vflight/commons/database"
	"github.com/vk-ai/vflight/commons/logger"
	"github.com/vk-ai/vflight/user-service/routes"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	log := logger.InitLogger("development")
	defer log.Sync()

	logger.Info("Starting User service", zap.String("service", "user-service"))

	// Initialize database config
	dbConfig := cfg.NewDatabaseConfig()

	// Create database connection
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		logger.Fatal("Failed to ping database", zap.Error(err))
	}

	logger.Info("Successfully connected to database",
		zap.String("host", dbConfig.Host),
		zap.String("database", dbConfig.DBName),
	)

	// Initialize router with routes
	router := routes.SetupRouter(db)

	// Start the server
	serverAddr := ":8081" // or get from config
	logger.Info("User service is starting", zap.String("address", serverAddr))

	if err := router.Run(serverAddr); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
