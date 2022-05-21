package DailyPrice

type CoinMarketStatus struct {
}

type CoinMarketQuote struct {
	USD CoinMarketPrice `json:"USD"`
}

type CoinMarketPrice struct {
	Price float64 `json:"price"`
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
