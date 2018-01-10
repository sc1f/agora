package main

import (
	"fmt"
	"os"
	"encoding/json"
	"flag"
	"strconv"
	"github.com/boltdb/bolt"
	"github.com/sc1f/agora/common"
	"github.com/sc1f/agora/api"
	"github.com/sc1f/agora/database"
)


func main() {
	// set CLI
	set_values := flag.Bool("set", false, "Set new values for BTC/ETH portfolio, example: --set ETH 0.123")
	historical := flag.Bool("h", false, "show historical values from the DB with a currency type, example: -h BTC")
	show_btc := flag.Bool("btc", false, "show BTC price")
	show_eth := flag.Bool("eth", false, "show ETH price")
	flag.Parse()

	if !*show_btc && !*show_eth && !*historical && !*set_values {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// set up our database
	db, err := database.SetupDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	if *set_values {
		// sets a new holding value
		currency_type := common.ValidateCurrencyType(flag.Arg(0))
		if currency_type != "BTC" && currency_type != "ETH" {
			panic("currency type not supported")
		}

		new_value, err := strconv.ParseFloat(flag.Arg(1), 64)
		if err != nil {
			panic(err)
		}

		err = database.SetHolding(db, currency_type, common.Holding{ Value: new_value })
		if err != nil {
			panic(err)
		}

		fmt.Println("New holding of " + currency_type + " has been set. Fetching current price...")
	}

	if *historical {
		// Shows all past data from the DB
		fmt.Println("Showing historical data...")
		fmt.Println("--------------------------")

		currency_type := common.ValidateCurrencyType(flag.Arg(0))

		// TODO: let's not have raw db operations in our main function?
		err := db.View(func(tx *bolt.Tx) error {
				bucket_key := []byte("PortfolioValue")
				values := tx.Bucket(bucket_key).Cursor()

				for key, val := values.First(); key != nil; key, val = values.Next() {
					var portfolio_value common.PortfolioValue
					json.Unmarshal(val, &portfolio_value)
					if portfolio_value.Type == currency_type {
						formatted_val := "$" + strconv.FormatFloat(portfolio_value.Value, 'f', 2, 64)
						fmt.Printf("type: %s | date: %s | holding: %s\n", currency_type, key, formatted_val)
					}
				}

				return nil
			})

		if err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	if *show_btc {
		btc_holding, err := database.GetHolding(db, "BTC")
		if err != nil {
			panic(err)
		}
		btc_price := api.GetPrice("BTC", 0).USD // current price
		portfolio_float := btc_price * btc_holding
		portfolio_value := strconv.FormatFloat(portfolio_float, 'f', 2, 64)
		fmt.Println("$" + string(portfolio_value))

		new_portfolio_value := common.PortfolioValue{ Type: "BTC", USD: btc_holding, Value: portfolio_float}
		err = database.AddPortfolioValue(db, new_portfolio_value)
		if err != nil {
			panic(err)
		}
	}
	
	if *show_eth {
		eth_holding, err := database.GetHolding(db, "ETH")
		if err != nil {
			panic(err)
		}
		var eth_price float64 = api.GetPrice("ETH", 0).USD // current price
		portfolio_float := eth_price * eth_holding
		portfolio_value := strconv.FormatFloat(portfolio_float, 'f', 2, 64)
		fmt.Println("$" + string(portfolio_value))

		new_portfolio_value := common.PortfolioValue{ Type: "ETH", USD: eth_holding, Value: portfolio_float}
		err = database.AddPortfolioValue(db, new_portfolio_value)
		if err != nil {
			panic(err)
		}
	}

	os.Exit(0)
}