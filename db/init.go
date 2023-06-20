package repo

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Transaction struct {
	Id int64 `json:"id"`
	Desc string `json:"desc"`
	Sum int64 `json:"sum"`
}

var db *sql.DB

func init() {
	connection, err := sql.Open("sqlite", "file:app.db"); 
	if err != nil {
		fmt.Println(err)
	}

	db = connection


	if err := migrate(db); err != nil {
		fmt.Println(err)
	}

	if err := seed(db); err != nil {
		fmt.Println(err)
	}

	// if err := db.Close(); err != nil {
	// 	fmt.Println(err)
	// }
}

func migrate(db *sql.DB) error {
	if _, err := db.Exec(`
	drop table if exists transactions;
	CREATE TABLE transactions
		   (id INTEGER PRIMARY KEY,
			desc TEXT DEFAULT '',
		   sum INTEGER);
	`); err != nil {
		return err
	}
	return nil
}

func seed(db *sql.DB) error {
	query := `
	INSERT INTO transactions (desc, sum) 
	values
	('one', 170), 
	('two', 30),
	('tree', 120);
	`;
	if _, err := db.Exec(query); err != nil {
		return err
	}
	return nil
}