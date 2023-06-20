package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)


type Transaction struct {
	Id int64
	Desc string
	Sum int64
}

func main() {
	db, _ := sql.Open("sqlite", "file:app.db")

	rows, _ := db.Query("select * from transactions")
	
	if _, err := db.Exec(`
		drop table if exists transactions;
		CREATE TABLE transactions
   		    (id INTEGER PRIMARY KEY,
   		     desc TEXT DEFAULT '',
   			sum INTEGER);
		`); err != nil {
			fmt.Println(err)
		}

	var transactions []Transaction
	for rows.Next(){
		var transaction Transaction
		if err := rows.Scan(&transaction.Id, &transaction.Desc, &transaction.Sum); err != nil {
			fmt.Println(err)
		}
		transactions = append(transactions, transaction)
	}

	fmt.Printf("%v", transactions)

	if err := db.Close(); err != nil {
		fmt.Println(err)
	}

}