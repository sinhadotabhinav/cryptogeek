package main

import (
  "fmt"
  "github.com/sinhadotabhinav/cryptogeek/internal/assets"
  "github.com/sinhadotabhinav/cryptogeek/internal/inputs"
  "github.com/sinhadotabhinav/cryptogeek/pkg/logger"
  "strings"
)

var (
	configFile = ConfigPath()
	config     = NewConfig(configFile)
  log        = logger.Logger()
)

func main() {
  log.Info("Application has started")
  fmt.Println("Welcome to cryptogeek application!")
  // calling exchange info api to fetch all live assets
  bases, quoteInput := assets.ExchangeAssets()
  // fetch base asset(s) as input from the user and retrieve live price details
  fmt.Printf("Enter one or multiple base assets seperated by comma from the below list:\n%s\n", bases)
  // interactive menu for users
  baseInput := inputs.UserInput()
  for counter := 0; counter < len(baseInput); counter++ {
    if !assets.AssetFound(bases, baseInput[counter]) {
      log.Fatalf("Invalid base asset entered: %s", baseInput[counter])
    }
    symbol := strings.ToUpper(baseInput[counter] + quoteInput)
    // print price details on the console
    fmt.Printf("\n--------------%s--------------\n", symbol)
    for key, value := range assets.PriceDetails(symbol) {
      fmt.Printf("%s: %s\n", key, value)
    }
    // if user has defined config file and purchase details then print profit/ loss statement
    if HasInvested(config) {
      present, key := BaseAssetPresent(config, baseInput[counter])
      if present {
        baseAsset := BaseAsset(config, key)
        fmt.Printf("total_profit: %f %s\n",
          assets.TotalProfit(BaseAssetQuantity(baseAsset), BaseAssetPrice(baseAsset), assets.PriceDetails(symbol)["current_price"]),
          strings.ToUpper(quoteInput))
      } else {
        log.Debugf("Purchase details are not found in the config file for asset entered by user: %s", symbol)
      }
    }
  }
}
