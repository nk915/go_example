package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/proullon/ramsql/driver"
	dbcontext "local-testing.com/nk915/database"
)

const (
	host     = ""
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = ""
)

func main() {
	fmt.Println("Hello World..")

	isMemDB := true

	var db *sql.DB
	{
		var err error
		var psqlconn string
		if isMemDB {
			db, err = sql.Open("ramsql", "TestInMemDB")
		} else {
			psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
			db, err = sql.Open("postgres", psqlconn)
		}

		if err != nil {
			fmt.Println("Err DB Conn: ", err)
			os.Exit(-1)
		}
	}

	repostiory, err := dbcontext.New(db)
	defer repostiory.Close()

	if err != nil {
		fmt.Println("Err New: ", err)
		os.Exit(-1)
	}

	repostiory.CreateTable()

	user := dbcontext.User{"kng", 30, "nk915@local-testing.com", "seoul", "010-1111-2222"}
	repostiory.InsertUser(user)
	repostiory.GetUserByID("kng")
	repostiory.GetUserByID("A")

	defer fmt.Println("main end")
}
