// package api interfaces with CryptoCompare
/* Package api makes requests to the CryptoCompare API. */

package api

import (
    "io/ioutil"
	"net/http"
	"encoding/json"
)

type Price struct {
	// represents a crypto price filled with data from the API
	USD float64 `json:"USD"`
}

type PortfolioValue struct {
	// represents the mapping of a crypto price to a portfolio holding, i.e holding * price
	USD float64
	crypto float64
}

// GetPrice interfaces with the CryptoCompare API and gets the current price of a currency, returning a new Price
func GetPrice(currency_type string) Price {
	
	if currency_type != "BTC" && currency_type != "ETH" {
		// only supports BTC and ETH at the moment
		panic("error: invalid request currency type")
	}

	request_url := "https://min-api.cryptocompare.com/data/price?fsym=" + currency_type + "&tsyms=USD"
	response, err := http.Get(request_url)
	if err != nil {
		panic(err.Error())
	} else {
		response_data, _ := ioutil.ReadAll(response.Body)
		var current_price Price // new(Price)
		json.Unmarshal(response_data, &price)
        return current_price
	}
}

func GetDayAveragePrice(currency_type string) Price {
	if currency_type != "BTC" && currency_type != "ETH" {
		// only supports BTC and ETH at the moment
		panic("error: invalid request currency type")
	}

	request_url := "https://min-api.cryptocompare.com/data/price?fsym=" + currency_type + "&tsyms=USD"
	response, err := http.Get(request_url)
	if err != nil {
		panic(err.Error())
	} else {
		response_data, _ := ioutil.ReadAll(response.Body)
		var current_price Price // new(Price)
		json.Unmarshal(response_data, &price)
        return current_price
	}
}

// GetPortfolioValue gets the price and current portfolio holding for a currency, returning a PortfolioValue
func GetPortfolioValue(currency_type string) {
	if currency_type != "BTC" && currency_type != "ETH" {
		// only supports BTC and ETH at the moment
		panic("error: invalid request currency type")
	}

	current_price := GetPrice(currency_type).USD
	portfolio_holding := 0.20301
	crypto_price := portfolio_holding * current_price 

	var portfolio_value = new(PortfolioValue{
			USD: current_price,
			crypto: crypto_price
		})

	return portfolio_value
}