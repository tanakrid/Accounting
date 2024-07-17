package transaction

import (
	"fmt"
)

type Record struct{
	IsExpense bool
	Amount float64
	TypeName string
	Description string
	Date string
}

var records map[string]Record = map[string]Record{}

func Add(k string, r Record) {
	fmt.Println("add transaction")
	_, ok := records[k]
	if !ok {
		records[k] = r
	} else {
		fmt.Printf("record'id %s has in database", k)
	}
}

func Show() map[string]Record {
	fmt.Println("show transaction")
	return records
}

func Del(k string) {
	fmt.Println("Del transaction")
	delete(records, k)
}

func Edit(k string, r Record) {
	fmt.Println("Edit transaction")
	_, ok := records[k]
	if ok {
		records[k] = r
	} else {
		fmt.Printf("record'id %s has not in database", k)
	}
}