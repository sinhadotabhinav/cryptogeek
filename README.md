# Cryptogeek

[![Go Report Card](https://goreportcard.com/badge/github.com/sinhadotabhinav/cryptogeek)](https://goreportcard.com/report/github.com/sinhadotabhinav/cryptogeek)

This is a golang-based application that retrieves live exchange information of the cryptocurrency market. Additionally, the application also calculates your profits or losses on preferred cryptocurrencies.

## Configurations

Update the full path of `CRYPTOGEEK_CONFIG_PATH` in [`config.go`](https://github.com/sinhadotabhinav/cryptogeek/blob/master/cmd/cryptogeek/config.go#L26).

Customise the [`/configs/config.json`](https://github.com/sinhadotabhinav/cryptogeek/blob/master/configs/config.json) based on your trading currencies. Set `"has_invested": false` if you are not interested in the profit/loss statement of your assets.
 See below example:

```
{
	"has_invested": true,
	"platform": "binance",
	"quote_asset_symbol": "USDT",
	"base_assets": [
		{
			"symbol": "DOGE",
			"price": 0.52399,
			"quantity": 491.80770
		},
		{
			"symbol": "XRP",
			"price": 1.4809,
			"quantity": 26.14383
		}
	]
}
```

## How to run the application?

Fork this repository and follow the [Configurations](#configurations) section to update the config file. Update the go module name based on your needs and push changes to GitHub. Replace the module name in the below command and execute application.

```
$ go get github.com/sinhadotabhinav/cryptogeek
```

Command crytogeek covers:
>Fetching live exchange information from exchange platform such as Binance\
Functionality asking user to input desired quote asset and base asset(s)\
Printing live price details for the assets\
Optionally, printing profit/ loss statement in case of trades

Run the `$ cryptogeek` command when inside `GOBIN` path or from anywhere after appending the `GOBIN` path to bash `PATH`.
