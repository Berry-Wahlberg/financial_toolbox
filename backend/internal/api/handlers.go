package handlers

import (
	"financial_toolbox/internal/model"
	"financial_toolbox/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Financial Toolbox API is running",
	})
}

func GetStockData(c *gin.Context) {
	symbol := c.Param("symbol")
	data := service.GetMockStockData(symbol)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func CalculateIndicator(c *gin.Context) {
	var req model.IndicatorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	result, err := service.CalculateIndicator(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

func GetAvailableIndicators(c *gin.Context) {
	indicators := service.GetAvailableIndicators()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    indicators,
	})
}
