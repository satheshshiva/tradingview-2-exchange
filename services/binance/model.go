package binance

type binance struct {
	url       string
	apiKey    string
	apiSecret string
	prodEnv   bool
}

type newTradeRequest struct {
	Symbol    string  `url:"symbol"`
	Side      string  `url:"side"`
	OrderType string  `url:"type"`
	Qty       float32 `url:"quantity"`
	Timestamp string  `url:"timestamp"`
}
