package main

import (
  "fmt"
  "github.com/sinhadotabhinav/cryptogeek/internal/assets"
  "github.com/sinhadotabhinav/cryptogeek/internal/inputs"
  "github.com/sinhadotabhinav/cryptogeek/pkg/configs"
  "strings"
)

var logger = configs.Logger()

func main() {
  logger.Info("Application has started")
  fmt.Println("Welcome to cryptogeek application!")
  // calling exchange info api to fetch all live assets
  bases, quoteInput := assets.ExchangeAssets()
  // fetch base asset(s) as input from the user and retrieve live price details
  fmt.Printf("Enter one or multiple base assets seperated by comma from the below list:\n%s\n", bases)
  baseInput := inputs.UserInput()
  for counter := 0; counter < len(baseInput); counter++ {
    if !assets.AssetFound(bases, baseInput[counter]) {
      logger.Fatalf("Invalid base asset entered: %s", baseInput[counter])
    }
    symbol := strings.ToUpper(baseInput[counter] + quoteInput)
    // print price details on the console
    fmt.Printf("-------------%s-------------\n", symbol)
    for key, value := range assets.PriceDetails(symbol) {
      fmt.Printf("%s: %s\n", key, value)
    }
  }
}
