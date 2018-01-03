package main

import (
	"fmt"
	//"os"
	"flag"
	"strconv"
	"github.com/sc1f/agora/api"
	"github.com/sc1f/agora/db"
)


func main() {
	// set CLI
	set_values := flag.Bool("set", false, "Set new values for btc/eth portfolio, example: --set ETH 0.123")
	flag.Parse()
	if *set_values {
		// call for input here
		currency_type := flag.Arg(0)
		new_value, err := strconv.ParseFloat(flag.Arg(1), 64)
		if err != nil {
			panic(err.Error())
		}
		set := db.SetPortfolioValue(currency_type, new_value)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(set)
		//fmt.Printf("type: %s, new_val: %e", currency_type, new_value)
		//os.Exit(1)
	}
	var eth_price float64 = api.GetPrice("ETH").USD
	var eth float64 = 0.20134147
	portfolio := eth * eth_price
	formatted := strconv.FormatFloat(portfolio, 'f', 2, 64)
	fmt.Println("$" + formatted)
}