package assets

import (
  "fmt"
  "github.com/sinhadotabhinav/cryptogeek/internal/inputs"
  "github.com/sinhadotabhinav/cryptogeek/pkg/api"
  "github.com/sinhadotabhinav/cryptogeek/pkg/logger"
  "github.com/sinhadotabhinav/cryptogeek/pkg/mappers"
  "github.com/sinhadotabhinav/cryptogeek/pkg/models"
  "strings"
  "sort"
)

var log = logger.Logger()

func AssetFound(assets []string, input string) bool {
  for _, value := range assets {
    if strings.EqualFold(value, input) {
      return true
    }
  }
  return false
}

func baseAssets(info models.ExchangeInfo, quote string) []string {
  baseAssets := make(map[string]struct{})
  for counter := 0; counter < len(info.Symbols_); counter++ {
    if strings.EqualFold(models.QuoteAsset(info.Symbols_[counter]), quote) {
      baseAssets[models.BaseAsset(info.Symbols_[counter])] = struct{}{}
    }
  }
  return sortMap(baseAssets)
}

func ExchangeAssets() ([]string, string) {
  resp, err := api.ExchangeInfo()
  if err != nil {
    log.Fatalf("Http request has failed: %s", err.Error())
  }
  info := mappers.ExchangeInfoMapper(resp)
  quotes := quoteAssets(info)
  // interactive menu for users
  fmt.Printf("Enter a quote asset from the below list:\n%s\n", quotes)
  quoteInput := inputs.UserInput()[0]
  if !AssetFound(quotes, quoteInput) {
    log.Fatalf("Invalid quote asset entered: %s", quoteInput)
  }
  return baseAssets(info, quoteInput), quoteInput
}

func PriceDetails(symbol string) map[string]string {
  resp, err := api.Price24Hour(symbol)
  if err != nil {
    log.Fatalf("Http request has failed: %s", err.Error())
  }
  price24Hour := mappers.Price24HourMapper(resp)
  priceMap := make(map[string]string)
  priceMap["current_price"] = price24Hour.LastPrice_
  priceMap["weighted_average_price"] = price24Hour.WeightedAvgPrice_
  priceMap["lowest_24h_price"] = price24Hour.LowPrice_
  priceMap["highest_24h_price"] = price24Hour.HighPrice_
  return priceMap
}

func quoteAssets(info models.ExchangeInfo) []string {
  quoteAssets := make(map[string]struct{})
  for counter := 0; counter < len(info.Symbols_); counter++ {
    quoteAssets[models.QuoteAsset(info.Symbols_[counter])] = struct{}{}
  }
  return sortMap(quoteAssets)
}

func sortMap(assets map[string]struct{}) []string {
  keys := make([]string, 0, len(assets))
  for key := range assets {
    keys = append(keys, key)
  }
  sort.Strings(keys)
  return keys
}
