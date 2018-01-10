package common

import (
	"strings"
	)

func ValidateCurrencyType(in_type string) string {
	if currency_type := strings.ToUpper(in_type); currency_type == "BTC" || currency_type == "ETH" {
		return currency_type
	} else {
		panic("Unsupported currency type!")
	}
	return ""
}