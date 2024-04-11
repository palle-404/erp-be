package app

import (
	"github.com/gin-gonic/gin"
	"github.com/palle-404/erp-be/src/service"
)

func createRouter(apiLayer service.Layer) *gin.Engine {
	router := gin.Default()

	apiGroup := router.Group("/api")

	apiGroup.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "working!"})
	})

	return router
}
