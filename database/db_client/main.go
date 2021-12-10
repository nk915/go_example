package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/proullon/ramsql/driver"
	ramsqldb "local-testing.com/nk915/database"
)

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
