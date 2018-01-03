package db

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func LoadDB() {
	db, err := bolt.Open("~/agora/agora.db", 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func CreateNewDB() {
	// add initial buckets
	db.Update(func (tx *bolt.Tx) error {
			portfolio, portfolio_err := tx.CreateBucket([]byte("PortfolioBucket"))
			price, price_err := tx.CreateBucket([]byte("PriceBucket"))
			if portfolio_err != nil || price_err != nil {
				return fmt.Errorf("create bucket: %s, %s", portfolio_err, price_err)
			}
			return nil
		})
}

// AddPortfolioValueToDB takes a PortfolioValue struct, converts its float64 values to byte[], and adds it to BoltDB
func AddPortfolioValueToDB() {

}