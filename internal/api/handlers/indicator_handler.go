package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"financial_toolbox/internal/model"
	"financial_toolbox/internal/service"
)

// GetIndicator 处理指标请求的API接口
// 参数：
// - symbol: 股票代码
// - indicator: 指标类型（如macd）
func GetIndicator(c *gin.Context) {
	// 获取请求参数
	symbol := c.Query("symbol")
	indicator := c.Query("indicator")

	// 验证参数
	if symbol == "" {
		c.JSON(http.StatusBadRequest, model.AnalysisResponse{
			Success: false,
			Error:   "股票代码不能为空",
		})
		return
	}

	if indicator == "" {
		c.JSON(http.StatusBadRequest, model.AnalysisResponse{
			Success: false,
			Error:   "指标类型不能为空",
		})
		return
	}

	// 模拟K线数据（实际项目中应该从数据库或API获取）
	klineData := []model.KLine{
		{Open: 100.0, High: 105.0, Low: 98.0, Close: 102.0, Volume: 10000, Date: "2023-01-01"},
		{Open: 102.0, High: 108.0, Low: 100.0, Close: 106.0, Volume: 12000, Date: "2023-01-02"},
		{Open: 106.0, High: 110.0, Low: 104.0, Close: 108.0, Volume: 15000, Date: "2023-01-03"},
		{Open: 108.0, High: 112.0, Low: 107.0, Close: 110.0, Volume: 13000, Date: "2023-01-04"},
		{Open: 110.0, High: 115.0, Low: 109.0, Close: 113.0, Volume: 18000, Date: "2023-01-05"},
		{Open: 113.0, High: 116.0, Low: 111.0, Close: 114.0, Volume: 14000, Date: "2023-01-06"},
		{Open: 114.0, High: 118.0, Low: 112.0, Close: 117.0, Volume: 16000, Date: "2023-01-07"},
		{Open: 117.0, High: 120.0, Low: 115.0, Close: 119.0, Volume: 19000, Date: "2023-01-08"},
		{Open: 119.0, High: 122.0, Low: 118.0, Close: 121.0, Volume: 17000, Date: "2023-01-09"},
		{Open: 121.0, High: 125.0, Low: 120.0, Close: 123.0, Volume: 20000, Date: "2023-01-10"},
		{Open: 123.0, High: 126.0, Low: 122.0, Close: 124.0, Volume: 15000, Date: "2023-01-11"},
		{Open: 124.0, High: 128.0, Low: 123.0, Close: 126.0, Volume: 18000, Date: "2023-01-12"},
		{Open: 126.0, High: 130.0, Low: 125.0, Close: 128.0, Volume: 22000, Date: "2023-01-13"},
		{Open: 128.0, High: 132.0, Low: 127.0, Close: 131.0, Volume: 21000, Date: "2023-01-14"},
		{Open: 131.0, High: 135.0, Low: 130.0, Close: 133.0, Volume: 25000, Date: "2023-01-15"},
	}

	// 根据指标类型选择对应的Python脚本路径
	var scriptPath string
	switch indicator {
	case "macd":
		scriptPath = "scripts/ta_indicators/macd.py"
	case "rsi":
		scriptPath = "scripts/ta_indicators/rsi.py"
	default:
		c.JSON(http.StatusBadRequest, model.AnalysisResponse{
			Success: false,
			Error:   "不支持的指标类型",
		})
		return
	}

	// 调用服务层计算指标
	result, err := service.CalculateIndicator(klineData, scriptPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.AnalysisResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.AnalysisResponse{
		Success: true,
		Result:  result,
	})
}