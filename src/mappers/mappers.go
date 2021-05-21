package mappers

import (
  "encoding/json"
  "github.com/sinhadotabhinav/cryptogeek/src/configs"
  "io/ioutil"
  "net/http"
)

var logger = configs.Logger()

func ExchangeInfoMapper(response *http.Response) configs.ExchangeInfo {
  bodyBytes := responseBody(response)
  var responseObject configs.ExchangeInfo
  json.Unmarshal(bodyBytes, &responseObject)
  return responseObject
}

func Price24HourMapper(response *http.Response) configs.Price24Hour {
  bodyBytes := responseBody(response)
  var responseObject configs.Price24Hour
  json.Unmarshal(bodyBytes, &responseObject)
  return responseObject
}

func responseBody(response *http.Response) []byte {
  defer response.Body.Close()
  bodyBytes, err := ioutil.ReadAll(response.Body)
  if err != nil {
    logger.Fatalf("Error reading response: %s", err.Error())
  }
  return bodyBytes
}
