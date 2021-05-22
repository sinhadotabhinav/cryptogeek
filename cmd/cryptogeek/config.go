package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// JSONConfig is the model for config json file
type JSONConfig struct {
	HasInvested      bool              `json:"has_invested"`
	Platform         string            `json:"platform"`
	QuoteAssetSymbol string            `json:"quote_asset_symbol"`
	BaseAssets       []BaseAssetConfig `json:"base_assets"`
}

// BaseAssetConfig is the model for base asset object in config json file
type BaseAssetConfig struct {
	Symbol   string  `json:"symbol"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

// CRYPTOGEEK_CONFIG_PATH is path to config file
const CRYPTOGEEK_CONFIG_PATH = "../../configs/config.json"

// BaseAsset func returns the base asset details
func BaseAsset(config JSONConfig, key int) BaseAssetConfig {
	return config.BaseAssets[key]
}

// BaseAssets func returns the slice of base asset details
func BaseAssets(config JSONConfig) []BaseAssetConfig {
	return config.BaseAssets
}

// BaseAssetPresent func validates if value input by user is present; if present
// returns true and its position in the slice
func BaseAssetPresent(config JSONConfig, baseAsset string) (bool, int) {
	baseAssets := config.BaseAssets
	for counter := 0; counter < len(baseAssets); counter++ {
		if strings.EqualFold(baseAsset, baseAssets[counter].Symbol) {
			return true, counter
		}
	}
	return false, -1
}

// BaseAssetPrice func returns the bought price of the base asset
func BaseAssetPrice(baseAsset BaseAssetConfig) float64 {
	return baseAsset.Price
}

// BaseAssetQuantity func returns the bought quantity of the base asset
func BaseAssetQuantity(baseAsset BaseAssetConfig) float64 {
	return baseAsset.Quantity
}

// ConfigPath func initialises the absolute path of config file
func ConfigPath() string {
	absPath, err := filepath.Abs(CRYPTOGEEK_CONFIG_PATH)
	if err != nil {
		log.Fatal("'CRYPTOGEEK_CONFIG_PATH' is not defined")
		return ""
	}
	return absPath
}

// HasInvested func fetches the value of has_invested parameter
func HasInvested(config JSONConfig) bool {
	if config.HasInvested {
		return true
	}
	return false
}

// NewConfig func reads and unmarshals the config file
func NewConfig(file string) JSONConfig {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Tried to read config file %s, got %s", file, err.Error())
	}
	var conf JSONConfig
	if err = json.Unmarshal(raw, &conf); err != nil {
		log.Fatalf("Tried to unmarshal config file %s, got %s", file, err.Error())
	}
	return conf
}
