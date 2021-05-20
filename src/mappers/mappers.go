package mappers

type ExchangeInfo struct {
	Timezone_        string        `json:"timezone"`
	ServerTime_      int64         `json:"serverTime"`
	RateLimits_      []interface{} `json:"rateLimits"`
	ExchangeFilters_ []interface{} `json:"exchangeFilters"`
	Symbols_         []string      `json:"symbols"`
}

const acceptHeader = "application/json"
const binanceUrl = "https://api.binance.com/api/v3"

func AcceptHeader() string {
 return acceptHeader
}

func BinanceBaseUrl() string {
 return binanceUrl
}
