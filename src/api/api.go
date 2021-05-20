package api

import (
  "github.com/sinhadotabhinav/cryptogeek/configs"
  "net/http"
)

var logger = configs.Logger()

func ExchangeInfo() (*http.Response, error) {
  return httpRequest(configs.GetMethod(), configs.ExchangeInfoEndpoint())
}

func Price() (*http.Response, error) {
  return httpRequest(configs.GetMethod(), configs.PriceEndpoint())
}

func Price24hour() (*http.Response, error) {
  return httpRequest(configs.GetMethod(), configs.Price24hourEndpoint())
}

func httpRequest(method string, url string) (*http.Response, error) {
  client := &http.Client{}
  req, err := http.NewRequest(method, url, nil)
  if err != nil {
    logger.Fatalf("Invalid http request %s: %s", req, err.Error())
  }
  req.Header.Add("Accept", configs.AcceptHeader())
  return client.Do(req)
}
