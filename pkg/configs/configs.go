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

type Price24Hour struct {
  Symbol_                     string        `json:"symbol"`
  PriceChange_                string        `json:"priceChange"`
  PriceChangePercent_         string        `json:"priceChangePercent"`
  WeightedAvgPrice_           string        `json:"weightedAvgPrice"`
  PrevClosePrice_             string        `json:"prevClosePrice"`
  LastPrice_                  string        `json:"lastPrice"`
  LastQty_                    string        `json:"lastQty"`
  BidPrice_                   string        `json:"bidPrice"`
  BidQty_                     string        `json:"bidQty"`
  AskPrice_                   string        `json:"askPrice"`
  AskQty_                     string        `json:"askQty"`
  OpenPrice_                  string        `json:"openPrice"`
  HighPrice_                  string        `json:"highPrice"`
  LowPrice_                   string        `json:"lowPrice"`
  Volume_                     string        `json:"volume"`
  QuoteVolume_                string        `json:"quoteVolume"`
  OpenTime_                   string        `json:"openTime"`
  CloseTime_                  int           `json:"closeTime"`
  FirstId_                    int           `json:"firstId"`
  LastId_                     int           `json:"lastId"`
  Count_                      int           `json:"count"`
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
const price24HourEndpoint = "/ticker/24hr"
const symbolParam = "?symbol="

func AcceptHeader() string {
 return acceptHeader
}

func BaseAsset(info SymbolInfo) string {
  return info.BaseAsset_
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

func Price24HourEndpoint() string {
  return binanceBaseUrl() + price24HourEndpoint + symbolParameter()
}

func QuoteAsset(info SymbolInfo) string {
  return info.QuoteAsset_
}

func symbolParameter() string {
  return symbolParam
}

func SymbolName(info SymbolInfo) string {
  return info.Symbol_
}
