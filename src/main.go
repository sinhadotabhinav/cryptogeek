package main

import (
  "fmt"
  "github.com/sinhadotabhinav/cryptogeek/src/api"
  "github.com/sinhadotabhinav/cryptogeek/src/configs"
  "github.com/sinhadotabhinav/cryptogeek/src/mappers"
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
  fmt.Println(info.Timezone_)
  fmt.Println(info.ServerTime_)
  for counter := 0; counter < len(info.Symbols_); counter++ {
    fmt.Printf("Symbol %d: %s\n", counter + 1, configs.SymbolName(info.Symbols_[counter]))
  }
}
