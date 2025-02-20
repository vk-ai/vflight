package main

import (
	cfg "github.com/vk-ai/vflight/commons/config"
	"github.com/vk-ai/vflight/commons/database"
	"github.com/vk-ai/vflight/commons/logger"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	log := logger.InitLogger("development")
	defer log.Sync() // flushes buffer, if any

	logger.Info("Starting API service", zap.String("service", "api-service"))

	// Initialize database config
	dbConfig := cfg.NewDatabaseConfig()

	// Create database connection
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		logger.Fatal("Failed to ping database", zap.Error(err))
	}

	logger.Info("Successfully connected to database",
		zap.String("host", dbConfig.Host),
		zap.String("database", dbConfig.DBName),
	)

	// Initialize router and other components
	// ... rest of your application code

	logger.Info("API service is ready to handle requests")
}
