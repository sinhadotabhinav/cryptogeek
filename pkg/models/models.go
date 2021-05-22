package models

// ExchangeInfo is the json model for ExchangeInfo http response
type ExchangeInfo struct {
	Timezone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	RateLimits      []interface{} `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []SymbolInfo  `json:"symbols"`
}

// Price24Hour is the json model for Price24Hour http response
type Price24Hour struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           string `json:"openTime"`
	CloseTime          int    `json:"closeTime"`
	FirstId            int    `json:"firstId"`
	LastId             int    `json:"lastId"`
	Count              int    `json:"count"`
}

// SymbolInfo is the json model for Symbols json array in ExchangeInfo http response
type SymbolInfo struct {
	Symbol                     string        `json:"symbol"`
	Status                     string        `json:"status"`
	BaseAsset                  string        `json:"baseAsset"`
	BaseAssetPrecision         int           `json:"baseAssetPrecision"`
	QuoteAsset                 string        `json:"quoteAsset"`
	QuotePrecision             int           `json:"quotePrecision"`
	QuoteAssetPrecision        int           `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int           `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int           `json:"quoteCommissionPrecision"`
	OrderTypes                 []interface{} `json:"orderTypes"`
	IcebergAllowed             bool          `json:"icebergAllowed"`
	OcoAllowed                 bool          `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed bool          `json:"quoteOrderQtyMarketAllowed"`
	IsSpotTradingAllowed       bool          `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool          `json:"isMarginTradingAllowed"`
	Filters                    []interface{} `json:"filters"`
	Permissions                []string      `json:"permissions"`
}

// BaseAsset func returns the base asset string in SymbolInfo
func BaseAsset(info SymbolInfo) string {
	return info.BaseAsset
}

// QuoteAsset func returns the quote asset string in SymbolInfo
func QuoteAsset(info SymbolInfo) string {
	return info.QuoteAsset
}

// SymbolName func returns the symbol name string in SymbolInfo
func SymbolName(info SymbolInfo) string {
	return info.Symbol
}
