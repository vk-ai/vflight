package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vk-ai/vflight/commons/database"
)

func SetupRouter(db *database.Database) *gin.Engine {
	router := gin.Default()
	router.GET("/health", healthCheck)
	return router
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
