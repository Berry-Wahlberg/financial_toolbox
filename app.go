package main

import (
	"context"
	"encoding/json"
	"fmt"
	"financial_toolbox/internal/model"
	"financial_toolbox/internal/service"
	"math/rand/v2"
	"time"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	fmt.Println("Financial Toolbox started")
}

func (a *App) shutdown(ctx context.Context) {
	fmt.Println("Financial Toolbox shutting down")
}

func (a *App) GetStockData(symbol string) ([]model.KLine, error) {
	return service.GetMockStockData(symbol), nil
}

func (a *App) CalculateIndicator(reqJSON string) (string, error) {
	var req model.IndicatorRequest
	if err := json.Unmarshal([]byte(reqJSON), &req); err != nil {
		return "", err
	}

	result, err := service.CalculateIndicator(req)
	if err != nil {
		return "", err
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(resultJSON), nil
}

func (a *App) GetAvailableIndicators() ([]string, error) {
	return service.GetAvailableIndicators(), nil
}

func (a *App) ExecutePythonScript(scriptPath string, inputData string) (string, error) {
	return service.ExecutePythonScript(scriptPath, inputData)
}

func (a *App) GetAppVersion() string {
	return "1.0.0"
}

func (a *App) GetSystemInfo() map[string]interface{} {
	return map[string]interface{}{
		"version":    "1.0.0",
		"platform":   "Windows",
		"buildTime":  time.Now().Format("2006-01-02 15:04:05"),
		"author":     "Financial Toolbox",
	}
}
