package service

import (
	"financial_toolbox/internal/model"
	"financial_toolbox/internal/python"
	"math/rand/v2"
	"time"
)

func GetMockStockData(symbol string) []model.KLine {
	rng := rand.New(rand.NewPCG(1234, 5678))
	var data []model.KLine

	basePrice := 100.0
	for i := 0; i < 100; i++ {
		change := (rng.Float64() - 0.5) * 4
		basePrice += change
		data = append(data, model.KLine{
			Date:   time.Now().AddDate(0, 0, -100+i).Format("2006-01-02"),
			Open:   basePrice - 1 + rng.Float64()*2,
			High:   basePrice + rng.Float64()*2,
			Low:    basePrice - rng.Float64()*2,
			Close:  basePrice,
			Volume: int64(100000 + rng.Int64N(50000)),
		})
	}
	return data
}

func CalculateIndicator(req model.IndicatorRequest) (*model.IndicatorResult, error) {
	data := GetMockStockData(req.Symbol)
	var result []map[string]interface{}

	for _, kline := range data {
		item := map[string]interface{}{
			"date": kline.Date,
		}

		switch req.Indicator {
		case "ma":
			item["value"] = kline.Close + (rand.Float64() - 0.5) * 5
		case "rsi":
			item["value"] = 30 + rand.Float64() * 40
		case "macd":
			item["macd"] = rand.Float64() * 2 - 1
			item["signal"] = rand.Float64() * 2 - 1
			item["histogram"] = rand.Float64() * 2 - 1
		}

		result = append(result, item)
	}

	return &model.IndicatorResult{
		Indicator: req.Indicator,
		Data:      result,
	}, nil
}

func GetAvailableIndicators() []string {
	return []string{
		"ma",
		"ema",
		"rsi",
		"macd",
		"bollinger",
		"kdj",
		"cci",
	}
}

func ExecutePythonScript(scriptPath string, inputData string) (string, error) {
	pm := python.NewPythonManager()
	return pm.ExecuteScript(scriptPath, inputData)
}
