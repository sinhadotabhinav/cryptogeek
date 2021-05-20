package main

import (
  "fmt"
  // "encoding/json"
  // "io/ioutil"
  "github.com/sinhadotabhinav/cryptogeek/api"
  "github.com/sinhadotabhinav/cryptogeek/configs"
  "github.com/sinhadotabhinav/cryptogeek/src/mappers"
)

var logger = configs.Logger()

func main() {
  logger.Debug("Starting application...")
  fmt.Println("Welcome to cryptogeek application")
  logger.Info("Application has started")
  resp, err := api.ExchangeInfo()
  if err != nil {
    logger.Fatalf("%s request has failed: %s", configs.GetMethod(), err.Error())
  }
  fmt.Println(resp)
  fmt.Println(mappers.BinanceBaseUrl())
  //
  // defer resp.Body.Close()
  // // bodyBytes, err := ioutil.ReadAll(resp.Body)
  // // if err != nil {
  // //   logger.Fatalf("Error reading response: %s", err.Error())
  // // }
  // // return bodyBytes
  //
  // // fmt.Printf("response: \n%s", bodyBytes)
  // var responseObject configs.ExchangeInfo
  // json.Unmarshal(bodyBytes, &responseObject)
  // // fmt.Printf("API Response as struct:\n%+v\n", responseObject)
  // fmt.Println(responseObject.Timezone_)
  // fmt.Println(responseObject.ServerTime_)
  // // fmt.Println(responseObject.RateLimits_)
  // // fmt.Println(responseObject.ExchangeFilters_)
  // // fmt.Println(responseObject.Symbols_)
  // for counter := 0; counter < len(responseObject.Symbols_); counter++ {
  //   fmt.Printf("Symbol %d: %s\n", counter + 1, configs.SymbolName(responseObject.Symbols_[counter]))
  // }
}
