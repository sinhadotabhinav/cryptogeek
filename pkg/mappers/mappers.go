package mappers

import (
  "encoding/json"
  "github.com/sinhadotabhinav/cryptogeek/pkg/logger"
  "github.com/sinhadotabhinav/cryptogeek/pkg/models"
  "io/ioutil"
  "net/http"
)

var log = logger.Logger()

func ExchangeInfoMapper(response *http.Response) models.ExchangeInfo {
  bodyBytes := responseBody(response)
  var responseObject models.ExchangeInfo
  json.Unmarshal(bodyBytes, &responseObject)
  return responseObject
}

func Price24HourMapper(response *http.Response) models.Price24Hour {
  bodyBytes := responseBody(response)
  var responseObject models.Price24Hour
  json.Unmarshal(bodyBytes, &responseObject)
  return responseObject
}

func responseBody(response *http.Response) []byte {
  defer response.Body.Close()
  bodyBytes, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatalf("Error reading response: %s", err.Error())
  }
  return bodyBytes
}
