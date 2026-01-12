package model

type KLine struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

type IndicatorRequest struct {
	Symbol    string                 `json:"symbol"`
	Indicator string                 `json:"indicator"`
	Params    map[string]interface{} `json:"params"`
}

type IndicatorResult struct {
	Indicator string                 `json:"indicator"`
	Data      []map[string]interface{} `json:"data"`
}

type StockData struct {
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
}
