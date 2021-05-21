package main

import (
  "bufio"
  "fmt"
  "github.com/sinhadotabhinav/cryptogeek/pkg/api"
  "github.com/sinhadotabhinav/cryptogeek/pkg/configs"
  "github.com/sinhadotabhinav/cryptogeek/pkg/mappers"
  "os"
  "strings"
  "sort"
)

var logger = configs.Logger()

func main() {
  logger.Info("Application has started")
  fmt.Println("Welcome to cryptogeek application!")
  // calling exchange info api to fetch all live assets
  bases, quoteInput := exchangeAssets()
  // fetch base asset(s) as input from the user and retrieve live price details
  fmt.Printf("Enter one or multiple base assets seperated by comma from the below list:\n%s\n", bases)
  baseInput := userInput()
  for counter := 0; counter < len(baseInput); counter++ {
    if !assetFound(bases, baseInput[counter]) {
      logger.Fatalf("Invalid base asset entered: %s", baseInput[counter])
    }
    symbol := strings.ToUpper(baseInput[counter] + quoteInput)
    // print price details on the console
    fmt.Printf("-------------%s-------------\n", symbol)
    for key, value := range priceDetails(symbol) {
      fmt.Printf("%s: %s\n", key, value)
    }
  }
}

func assetFound(assets []string, input string) bool {
  for _, value := range assets {
    if strings.EqualFold(value, input) {
      return true
    }
  }
  return false
}

func assetPrice(baseInput []string, quoteInput string) string {
  fmt.Println(baseInput)
  fmt.Println(quoteInput)
  return "OK"
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

func exchangeAssets() ([]string, string) {
  resp, err := api.ExchangeInfo()
  if err != nil {
    logger.Fatalf("Http request has failed: %s", err.Error())
  }
  info := mappers.ExchangeInfoMapper(resp)
  quotes := quoteAssets(info)
  // interactive menu for users
  fmt.Printf("Enter a quote asset from the below list:\n%s\n", quotes)
  quoteInput := userInput()[0]
  if !assetFound(quotes, quoteInput) {
    logger.Fatalf("Invalid quote asset entered: %s", quoteInput)
  }
  return baseAssets(info, quoteInput), quoteInput
}

func priceDetails(symbol string) map[string]string {
  resp, err := api.Price24Hour(symbol)
  if err != nil {
    logger.Fatalf("Http request has failed: %s", err.Error())
  }
  price24Hour := mappers.Price24HourMapper(resp)
  priceMap := make(map[string]string)
  priceMap["current_price"] = price24Hour.LastPrice_
  priceMap["weighted_average_price"] = price24Hour.WeightedAvgPrice_
  priceMap["lowest_24h_price"] = price24Hour.LowPrice_
  priceMap["highest_24h_price"] = price24Hour.HighPrice_
  return priceMap
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

func userInput() []string {
  // this function removes the additional \n character at the end of each
  // line input added by bufio. this also removes whitespaces if any
  text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  return strings.Split(strings.ReplaceAll(text[:len(text) - 1], " ", ""), ",")
}
