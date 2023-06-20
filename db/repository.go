package repo

import (
	"fmt"
)

func GetList() []Transaction {
	rows, _ := db.Query("select * from transactions")
	var transactions []Transaction
	for rows.Next(){
		var transaction Transaction
		if err := rows.Scan(&transaction.Id, &transaction.Desc, &transaction.Sum); err != nil {
			fmt.Println(err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions
}

func GetOne(id int) (Transaction, error) {
	var t Transaction
	row := db.QueryRow("SELECT * from transactions WHERE id = ?", id)
	if err := row.Scan(&t.Id, &t.Desc, &t.Sum); err != nil {
		return t, err
	}

	return t, nil
}

func Save(t Transaction) (int64, error) {
	result, _ := db.Exec("INSERT INTO transactions (desc, sum) values (?, ?)", t.Desc, t.Sum)
	id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    return id, nil
}

func Delete(id int) error {
	result, err := db.Exec("DELETE FROM transactions WHERE id = ?", id)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if count == 0 {
		return fmt.Errorf("not found value with id %d", id)
	}

	return err
}