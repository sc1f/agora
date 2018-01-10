package database

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/sc1f/agora/common"
)
 
// SetHolding...
func SetHolding(db *bolt.DB, currency_type string, holding common.Holding) error {
	holding_bytes, err := json.Marshal(holding)
	if err != nil {
		return fmt.Errorf("Could not marshal holding JSON: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("Holding")).Put([]byte(currency_type), holding_bytes)
		if err != nil {
			return fmt.Errorf("Could not set holding: %v", err)
		}
		return nil
	})
	fmt.Println("Your new holding has been set!")
	return nil
}

// GetHolding...
func GetHolding(db *bolt.DB, currency_type string) (float64, error) {
	var holding_bytes []byte
	
	err := db.View(func(tx *bolt.Tx) error {
			currency_key := []byte(currency_type)
			bucket_key := []byte("Holding")
			holding := tx.Bucket(bucket_key).Get(currency_key)
			holding_bytes = make([]byte, len(holding))
			copy(holding_bytes, holding)
			return nil
		})

	if err != nil {
		panic(err)
	}

	var current_holding common.Holding
	json.Unmarshal(holding_bytes, &current_holding)
	return current_holding.Value, nil
} 