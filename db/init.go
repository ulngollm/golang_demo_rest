package repo

import (
	"fmt"

	"github.com/go-faker/faker/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type Transaction struct {
	gorm.Model
	Desc string `json:"desc"`
	Sum int64 `json:"sum"`
}

var db *gorm.DB

func init() {
	connection, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{}); 
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

}

func migrate(db *gorm.DB) error {
	db.AutoMigrate(&Transaction{})
	return nil
}

func seed(db *gorm.DB) error {
	var transaction []Transaction
	for i := 0; i < 10; i++ {
		t := Transaction{
			Desc: faker.Word(),
			Sum: 100,
		}
		transaction = append(transaction, t )
	} 

	db.CreateInBatches(&Transaction{}, 4)
	return nil
}