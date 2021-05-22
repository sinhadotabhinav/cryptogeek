package assets

import (
	"fmt"
	"github.com/sinhadotabhinav/cryptogeek/internal/inputs"
	"github.com/sinhadotabhinav/cryptogeek/pkg/api"
	"github.com/sinhadotabhinav/cryptogeek/pkg/logger"
	"github.com/sinhadotabhinav/cryptogeek/pkg/mappers"
	"github.com/sinhadotabhinav/cryptogeek/pkg/models"
	"sort"
	"strconv"
	"strings"
)

var log = logger.Logger()

// AssetFound func validates if input by user is present in the exchange list
func AssetFound(assets []string, input string) bool {
	for _, value := range assets {
		if strings.EqualFold(value, input) {
			return true
		}
	}
	return false
}

func baseAssets(info models.ExchangeInfo, quote string) []string {
	baseAssets := make(map[string]struct{})
	for counter := 0; counter < len(info.Symbols); counter++ {
		if strings.EqualFold(models.QuoteAsset(info.Symbols[counter]), quote) {
			baseAssets[models.BaseAsset(info.Symbols[counter])] = struct{}{}
		}
	}
	return sortMap(baseAssets)
}

// ExchangeAssets func retrieves quote assets and base assets
func ExchangeAssets() ([]string, string) {
	resp, err := api.ExchangeInfo()
	if err != nil {
		log.Fatalf("Http request has failed: %s", err.Error())
	}
	info := mappers.ExchangeInfoMapper(resp)
	quotes := quoteAssets(info)
	fmt.Printf("Enter a quote asset from the below list:\n%s\n", quotes)
	quoteInput := inputs.UserInput()[0]
	if !AssetFound(quotes, quoteInput) {
		log.Fatalf("Invalid quote asset entered: %s", quoteInput)
	}
	return baseAssets(info, quoteInput), quoteInput
}

// PriceDetails returns a price details map of a base asset
func PriceDetails(symbol string) map[string]string {
	resp, err := api.Price24Hour(symbol)
	if err != nil {
		log.Fatalf("Http request has failed: %s", err.Error())
	}
	price24Hour := mappers.Price24HourMapper(resp)
	priceMap := make(map[string]string)
	priceMap["current_price"] = price24Hour.LastPrice
	priceMap["weighted_average_price"] = price24Hour.WeightedAvgPrice
	priceMap["lowest_24h_price"] = price24Hour.LowPrice
	priceMap["highest_24h_price"] = price24Hour.HighPrice
	return priceMap
}

func quoteAssets(info models.ExchangeInfo) []string {
	quoteAssets := make(map[string]struct{})
	for counter := 0; counter < len(info.Symbols); counter++ {
		quoteAssets[models.QuoteAsset(info.Symbols[counter])] = struct{}{}
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

// TotalProfit calculates profit/loss statement for a given asset
func TotalProfit(quantity float64, boughtPrice float64, currentPrice string) float64 {
	cPrice, _ := strconv.ParseFloat(currentPrice, 64)
	return quantity * (cPrice - boughtPrice)
}
