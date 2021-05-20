package configs

import (
  "github.com/sirupsen/logrus"
  "os"
)

type ExchangeInfo struct {
	Timezone_        string        `json:"timezone"`
	ServerTime_      int64         `json:"serverTime"`
	RateLimits_      []interface{} `json:"rateLimits"`
	ExchangeFilters_ []interface{} `json:"exchangeFilters"`
	Symbols_         []SymbolInfo  `json:"symbols"`
}

type SymbolInfo struct {
  Symbol_                     string        `json:"symbol"`
  Status_                     string        `json:"status"`
  BaseAsset_                  string        `json:"baseAsset"`
  BaseAssetPrecision_         int           `json:"baseAssetPrecision"`
  QuoteAsset_                 string        `json:"quoteAsset"`
  QuotePrecision_             int           `json:"quotePrecision"`
  QuoteAssetPrecision_        int           `json:"quoteAssetPrecision"`
  BaseCommissionPrecision_    int           `json:"baseCommissionPrecision"`
  QuoteCommissionPrecision_   int           `json:"quoteCommissionPrecision"`
  OrderTypes_                 []interface{} `json:"orderTypes"`
  IcebergAllowed_             bool          `json:"icebergAllowed"`
  OcoAllowed_                 bool          `json:"ocoAllowed"`
  QuoteOrderQtyMarketAllowed_ bool          `json:"quoteOrderQtyMarketAllowed"`
  IsSpotTradingAllowed_       bool          `json:"isSpotTradingAllowed"`
  IsMarginTradingAllowed_     bool          `json:"isMarginTradingAllowed"`
  Filters_                    []interface{} `json:"filters"`
  Permissions_                []string      `json:"permissions"`
}

const acceptHeader = "application/json"
const binanceUrl = "https://api.binance.com/api/v3"
const exchangeInfoEndpoint = "/exchangeInfo"
const getMethod = "GET"
const priceEndpoint = "/ticker/price"
const price24hourEndpoint = "/ticker/24hr"
const symbolParam = "?symbol="

func AcceptHeader() string {
 return acceptHeader
}

func binanceBaseUrl() string {
 return binanceUrl
}

func ExchangeInfoEndpoint() string {
 return binanceBaseUrl() + exchangeInfoEndpoint
}

func GetMethod() string {
 return getMethod
}

func Logger() *logrus.Logger {
  var log = logrus.New()
  log.Out = os.Stdout
  log.Formatter = &logrus.JSONFormatter{}
  log.Level = logrus.DebugLevel
  return log
}

func PriceEndpoint() string {
  return binanceBaseUrl() + priceEndpoint + symbolParameter()
}

func Price24hourEndpoint() string {
  return binanceBaseUrl() + price24hourEndpoint + symbolParameter()
}

func symbolParameter() string {
  return symbolParam
}

func SymbolName(info SymbolInfo) string {
  return info.Symbol_
}
