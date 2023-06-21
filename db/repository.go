package repo

func GetList() ([]Transaction, error) {
	var transactions []Transaction
	result := db.Find(&transactions)

	return transactions, result.Error
}

func GetOne(id int) (Transaction, error) {
	var t Transaction
	result := db.First(t, id)
	return t, result.Error
}

func Save(t Transaction) (Transaction, error) {
	result := db.Create(&t)
    return t, result.Error
}

func Delete(id int) error {
	result:= db.Delete(&Transaction{}, id)
	// if not found ?
	
	return result.Error
}