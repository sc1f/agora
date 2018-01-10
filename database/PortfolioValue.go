package database

import (
	"fmt"
	"time"
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/sc1f/agora/common"
)

// AddPortfolioValue takes a PortfolioValue struct, converts its float64 values to byte[], and adds it to BoltDB
func AddPortfolioValue(db *bolt.DB, portfolio_value common.PortfolioValue) error {
	portfolio_value_bytes, err := json.Marshal(portfolio_value)
	if err != nil {
		return fmt.Errorf("Could not marshal PortfolioValue: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
			err := tx.Bucket([]byte("PortfolioValue")).Put([]byte(time.Now().Format(time.RFC3339)), portfolio_value_bytes)
			if err != nil {
				return fmt.Errorf("Could not insert PortfolioValue: %v", err)
			}
			return nil
		})
	fmt.Println("Latest PortfolioValue has been saved to DB.")
	return err
}