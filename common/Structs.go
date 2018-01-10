// package common holds common structs used throughout the program
package common

type Holding struct {
	// holding represents the user's holding of a crypto currency
	Value float64
}

type Price struct {
	// represents a crypto price filled with data from the API
	USD float64 `json:"USD"`
}

type PortfolioValue struct {
	// represents the mapping of a crypto price to a portfolio holding, i.e holding * price
	Type string
	USD float64
	Value float64
}