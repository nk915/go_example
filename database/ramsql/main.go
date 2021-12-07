package main

import (
	"database/sql"
	"fmt"
)

func LoadUserAddresses(db *sql.DB, userID int64) ([]string, error) {
	query := `SELECT address.street_number, address.street FROM address 
							JOIN user_addresses ON address.id=user_addresses.address_id 
							WHERE user_addresses.user_id = $1;`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	var addresses []string
	for rows.Next() {
		var number int
		var street string
		if err := rows.Scan(&number, &street); err != nil {
			return nil, err
		}
		addresses = append(addresses, fmt.Sprintf("%d %s", number, street))
	}

	return addresses, nil
}

func main() {
	fmt.Println("Hello World..")
}
