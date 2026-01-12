package api

import (
	"financial_toolbox/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/stock/:symbol", handlers.GetStockData)
		api.POST("/indicator/calculate", handlers.CalculateIndicator)
		api.GET("/indicators", handlers.GetAvailableIndicators)
	}
}
