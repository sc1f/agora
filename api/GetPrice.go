// package api interfaces with CryptoCompare
package api

import (
    "io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/sc1f/agora/common"
)

// GetPrice interfaces with the CryptoCompare API and gets the current price of a currency at a given timestamp (optional), returning a new Price
func GetPrice(currency_type string, timestamp int) common.Price {

	if currency_type != "BTC" && currency_type != "ETH" {
		// only supports BTC and ETH at the moment
		panic("error: invalid request currency type")
	}

	request_type := "price"
	ts := ""

	if timestamp > 0 {
		// get a price at a timestamp
		request_type = "pricehistorical"
		ts = "&ts=" + string(timestamp)
	}

	request_url := "https://min-api.cryptocompare.com/data/" + request_type +"?fsym=" + currency_type + "&tsyms=USD" + ts
	response, err := http.Get(request_url)
	if err != nil {
		panic(err.Error())
	} else {
		response_data, _ := ioutil.ReadAll(response.Body)
		var current_price common.Price
		json.Unmarshal(response_data, &current_price)
        return current_price
	}
}