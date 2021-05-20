package main

import (
  "bufio"
  "fmt"
  "github.com/sinhadotabhinav/cryptogeek/src/api"
  "github.com/sinhadotabhinav/cryptogeek/src/configs"
  "github.com/sinhadotabhinav/cryptogeek/src/mappers"
  "os"
  "strings"
  "sort"
)

var logger = configs.Logger()

func main() {
  logger.Debug("Starting application...")
  fmt.Println("Welcome to cryptogeek application")
  logger.Info("Application has started")
  resp, err := api.ExchangeInfo()
  if err != nil {
    logger.Fatalf("Http request has failed: %s", err.Error())
  }
  info := mappers.ExchangeInfoMapper(resp)
  quotes := quoteAssets(info)
  // menu
  fmt.Printf("Enter a quote asset %s:\n", quotes)
  text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  input := text[:len(text)-1]
  if !assetFound(quotes, input) {
    logger.Fatalf("Invalid quote asset entered: %s", input)
  }
  bases := baseAssets(info, input)
  fmt.Printf("Enter one or multiple base assets seperated by comma %s:\n", bases)
}

func assetFound(assets []string, input string) bool {
  for _, value := range assets {
    if strings.EqualFold(value, input) {
      return true
    }
  }
  return false
}

func baseAssets(info configs.ExchangeInfo, quote string) []string {
  baseAssets := make(map[string]struct{})
  for counter := 0; counter < len(info.Symbols_); counter++ {
    if strings.EqualFold(configs.QuoteAsset(info.Symbols_[counter]), quote) {
      baseAssets[configs.BaseAsset(info.Symbols_[counter])] = struct{}{}
    }
  }
  return sortMap(baseAssets)
}

func quoteAssets(info configs.ExchangeInfo) []string {
  quoteAssets := make(map[string]struct{})
  for counter := 0; counter < len(info.Symbols_); counter++ {
    quoteAssets[configs.QuoteAsset(info.Symbols_[counter])] = struct{}{}
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
