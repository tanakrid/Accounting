package main

import (
	// "fmt"
	// "github.com/tanakrid/accounting/transaction"
	// "github.com/tanakrid/accounting/report"
	// "github.com/tanakrid/accounting/filter"
	"github.com/tanakrid/accounting/route"
)

func main() {
	// r1 := transaction.Record{
	// 	IsExpense: true, 
	// 	Amount: 100, 
	// 	TypeName: "food", 
	// 	Date: "2024/07/15 00:19",
	// }
	// r2 := transaction.Record{
	// 	IsExpense: true, 
	// 	Amount: 100, 
	// 	TypeName: "shopping", 
	// 	Date: "2024/07/16 00:19",
	// }
	// r3 := transaction.Record{
	// 	IsExpense: true, 
	// 	Amount: 100, 
	// 	TypeName: "learning", 
	// 	Date: "2024/07/17 00:19",
	// }
	// r4 := transaction.Record{
	// 	IsExpense: true, 
	// 	Amount: 100, 
	// 	TypeName: "food", 
	// 	Date: "2024/07/18 00:19",
	// }
	// transaction.Add("0001", r1)
	// transaction.Add("0002", r2)
	// transaction.Add("0003", r3)
	// transaction.Add("0004", r4)
	
	// fmt.Println("PieChart:> ", report.ShowPieChart(transaction.Show()))
	// fmt.Println("Filtered:> ", filter.ByType("food", transaction.Show()))
	// fmt.Println("Filtered:> ", filter.ByDate("2024/07/16 00:19", "2024/07/17 00:19", transaction.Show()))
	route.InitRoute()
}