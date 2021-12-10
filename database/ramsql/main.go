package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/proullon/ramsql/driver"
	ramsqldb "local-testing.com/nk915/database"
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

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("ramsql", "TestInMemDB")
		if err != nil {
			fmt.Println("Err ramsql", err)
			os.Exit(-1)
		}
	}

	repostiory, err := ramsqldb.New(db)
	defer repostiory.Close()

	if err != nil {
		fmt.Println("Err ramsql", err)
		os.Exit(-1)
	}

	repostiory.CreateTable()

	user := ramsqldb.User{"kng", 30, "nk915@local-testing.com", "seoul", "010-1111-2222"}
	repostiory.InsertUser(user)
	repostiory.GetUserByID("kng")
	repostiory.GetUserByID("A")

	defer fmt.Println("main end")
}
