package DailyPrice

type CoinMarketStatus struct {
}

type CoinMarketQuote struct {
	USD CoinMarketPrice `json:"USD"`
}

type CoinMarketPrice struct {
	Price     float64 `json:"price"`
	Volume24h float64 `json:"volume_24h"`
}

type CoinMarketData struct {
	Id    int             `json:"id"`
	Name  string          `json:"name"`
	Quote CoinMarketQuote `json:"quote"`
}

type CoinMarket struct {
	Status CoinMarketStatus `json:"status"`
	Data   []CoinMarketData `json:"data"`
}
