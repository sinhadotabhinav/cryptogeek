package api

import (
	"github.com/sinhadotabhinav/cryptogeek/pkg/logger"
	"net/http"
)

const jsonHeader = "application/json"
const binanceUrl = "https://api.binance.com/api/v3"
const exchangeEndpoint = "/exchangeInfo"
const getMethod = "GET"
const priceEndpoint = "/ticker/price"
const price24HourEndpoint = "/ticker/24hr"
const symbolParam = "?symbol="

var log = logger.Logger()

func acceptHeader() string {
	return jsonHeader
}

func binanceBaseUrl() string {
	return binanceUrl
}

// ExchangeInfo func performs http request to ExchangeInfo endpoint
func ExchangeInfo() (*http.Response, error) {
	return httpRequest(httpMethod(), exchangeInfoEndpoint())
}

func exchangeInfoEndpoint() string {
	return binanceBaseUrl() + exchangeEndpoint
}

func httpMethod() string {
	return getMethod
}

func httpRequest(method string, url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalf("Invalid http request %s: %s", req, err.Error())
	}
	req.Header.Add("Accept", acceptHeader())
	return client.Do(req)
}

// Price func performs http request to Price endpoint
func Price(symbol string) (*http.Response, error) {
	return httpRequest(httpMethod(), priceDetailsEndpoint()+symbol)
}

func priceDetailsEndpoint() string {
	return binanceBaseUrl() + priceEndpoint + symbolParameter()
}

// Price24Hour func performs http request to Price24Hour endpoint
func Price24Hour(symbol string) (*http.Response, error) {
	return httpRequest(httpMethod(), price24HourDetailsEndpoint()+symbol)
}

func price24HourDetailsEndpoint() string {
	return binanceBaseUrl() + price24HourEndpoint + symbolParameter()
}

func symbolParameter() string {
	return symbolParam
}
