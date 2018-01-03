package db

import (
	"fmt"
	//"github.com/boltdb/bolt"
)

// SetPortfolioValue saves the new into boltdb
func SetPortfolioValue(currency_type string, new_value float64) bool {
	fmt.Printf("type: %s, new_val: %e", currency_type, new_value)
	return true
}

func AddPortfolioValue(currency_type string, value_to_add float64) bool {
	return true
}

func SaveNewPrice(currency_type string, new_price float64) bool {
	return true 
}