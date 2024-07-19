package main

import (
	_ "github.com/proullon/ramsql/driver"
	"github.com/tanakrid/accounting/db"
	// "github.com/tanakrid/accounting/transaction"
	// "fmt"
	"github.com/tanakrid/accounting/route"
)

func main() {
	db.InitDB()
	// id := transaction.Add(transaction.Record{
	// 	IsExpense: true,
	// 	Amount: 100,
	// 	TypeName: "food",
	// 	Description: "noodle",
	// 	Date: "2024/07/19 10.52",
	// })
	// fmt.Println("created record's id:", id)
	// transaction.Show()
	route.InitRoute()
}