package database

import (
	"fmt"
	"github.com/boltdb/bolt"
)


// SetupDB creates and sets up our DB with buckets
func SetupDB() (*bolt.DB, error) {
	db, err := bolt.Open("agora.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not open db: %v", err)
	}
	// add initial buckets
	err = db.Update(func (tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte("Holding"))
			// error checking in go: return item of type err
			if err != nil {
				return fmt.Errorf("Could not create Holding bucket: %v", err)
			}
			// TODO: add separate buckets for ETH and BTC holdings
			_, err = tx.CreateBucketIfNotExists([]byte("PortfolioValue"))
			if err != nil {
				return fmt.Errorf("Could not create PortfolioValue bucket: %v", err)
			}
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("Could not create buckets: %v", err)
	}
	fmt.Println("DB access complete.")
	return db, nil
}