package main

import (
	"encoding/json"
	"io/ioutil"
  "path/filepath"
	"strings"
)

type JSONConfig struct {
	HasInvested_ 			 bool              `json:"has_invested"`
	Platform_    			 string            `json:"platform"`
	QuoteAssetSymbol_  string            `json:"quote_asset_symbol"`
	BaseAssets_  			 []BaseAssetConfig `json:"base_assets"`
}

type BaseAssetConfig struct {
  Symbol_   				 string        		 `json:"symbol"`
	Price_    	 			 float64           `json:"price"`
  Quantity_    		   float64           `json:"quantity"`
}

const CRYPTOGEEK_CONFIG_PATH = "../../configs/config.json"

func BaseAsset(config JSONConfig, key int) BaseAssetConfig {
	return config.BaseAssets_[key]
}

func BaseAssets(config JSONConfig) []BaseAssetConfig {
	return config.BaseAssets_
}

func BaseAssetPresent(config JSONConfig, baseAsset string) (bool, int) {
	baseAssets := config.BaseAssets_
	for counter := 0; counter < len(baseAssets); counter++ {
    if strings.EqualFold(baseAsset, baseAssets[counter].Symbol_) {
      return true, counter
    }
  }
	return false, -1
}


func BaseAssetPrice(baseAsset BaseAssetConfig) float64 {
	return baseAsset.Price_
}

func BaseAssetQuantity(baseAsset BaseAssetConfig) float64 {
	return baseAsset.Quantity_
}

func ConfigPath() string {
  absPath, err := filepath.Abs(CRYPTOGEEK_CONFIG_PATH)
	if err != nil {
    log.Fatal("'CRYPTOGEEK_CONFIG_PATH' is not defined")
		return ""
	}
	return absPath
}

func HasInvested(config JSONConfig) bool {
	if (config.HasInvested_) {
		return true
	}
	return false
}

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
