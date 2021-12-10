package dbcontext

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID      string
	Age     int64
	Email   string
	Address string
	Phone   string
}

type repository struct {
	db *sql.DB
}

type Repository interface {
	CreateTable() error
	InsertUser(user User) error
	DeleteUser(id string) error
	GetUserByID(id string) error
	Close() error
}

func New(db *sql.DB) (Repository, error) {
	return &repository{
		db: db,
	}, nil
}

func (repo *repository) CreateTable() error {
	tables := []string{
		//`CREATE TABLE user (id BIGSERIAL PRIMARY KEY, age INT, email TEXT, address TEXT, phone TEXT);`,
		//`DROP TABLE tbl_user;`,
		`CREATE TABLE IF NOT EXISTS tbl_user (id TEXT, age INT, email TEXT, address TEXT, phone TEXT);`,
		`CREATE TABLE IF NOT EXISTS tbl_user_addresses (address_id INT, user_id INT);`,
	}

	for _, table := range tables {
		_, err := repo.db.Exec(table)
		if err != nil {
			fmt.Println("Fail to err:", err)
			return err
		}
		fmt.Println("query_exec:", table)
	}
	return nil
}

func (repo *repository) InsertUser(user User) error {
	query := `INSERT INTO tbl_user (id, age, email, address, phone) VALUES ($1, $2, $3, $4, $5);`
	// query := `INSERT INTO user (id, age, email, address, phone) VALUES ('kng', 30, 'nk915@hanssak.co.kr', 'seoul', '010-1111-2222');`

	_, err := repo.db.Exec(query, user.ID, user.Age, user.Email, user.Address, user.Phone)
	//_, err := repo.db.Exec(query)
	if err != nil {
		fmt.Println("Err InsertUser:", err)
		return err
	}
	fmt.Println("query_exec:", query)
	return nil
}

func (repo *repository) DeleteUser(id string) error {
	return nil
}

func (repo *repository) GetUserByID(id string) error {
	var userRow = User{}
	query := `
		SELECT id, age, email, address, phone FROM tbl_user WHERE id = $1`

	if err := repo.db.QueryRow(query, id).
		Scan(&userRow.ID, &userRow.Age, &userRow.Email, &userRow.Address, &userRow.Phone); err != nil {
		fmt.Println("Err GetUserByID: ", err)
		return err
	}

	fmt.Println(userRow)
	return nil
}

func (repo *repository) Close() error {
	return repo.db.Close()
}
