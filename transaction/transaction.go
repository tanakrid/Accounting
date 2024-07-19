package transaction

import (
	"fmt"
	"log"
	"strconv"
	"github.com/tanakrid/accounting/db"
)

type Record struct {
    Id          int64
    IsExpense   bool
    Amount      float64
    TypeName    string
    Description string
    Date        string
}

func Add(r Record) int64{
	insert := `
	INSERT INTO transactions(is_expense, amount, type_name, description, date)
	VALUES (?,?,?,?,?);
	`
	stmt, err := db.DB.Prepare(insert)
	if err != nil {
		log.Fatal("prepare statement error:", err)
	}
	result, err := stmt.Exec(
		r.IsExpense, 
		r.Amount,
		r.TypeName,
		r.Description,
		r.Date,
	)
	if err != nil {
		log.Fatal("insert error:", err)
	}
	id, err := result.LastInsertId()
	fmt.Printf("record's id:%v, error:%v\n", id, err)
	rf, err := result.RowsAffected()
	fmt.Printf("RowsAffected:%v, error:%v\n", rf, err)
	return id
}

func Show() []Record {
	records := []Record{}
	rows, err := db.DB.Query(`
	SELECT id, is_expense, amount, type_name, description, date 
	FROM transactions`)
	if err != nil {
		log.Fatal("query error", err)
	}
	for rows.Next() {
		record := Record{}
		err := rows.Scan(&record.Id, &record.IsExpense, 
			&record.Amount, &record.TypeName, 
			&record.Description, &record.Date)
		if err != nil {
			log.Fatal("for rows error", err)
		}
		records = append(records, record)
	}
	return records
}

func ShowById(id string) Record {
	strId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("convert id error:%#v\n", err)
		log.Fatal("convert id error:", err)
	}
	row := db.DB.QueryRow(`
	SELECT id, is_expense, amount, type_name, description, date 
	FROM transactions
	WHERE id = ?`, strId)
	if err != nil {
		fmt.Printf("query error:%#v\n", err)
		log.Fatal("query error", err)
	}
	record := Record{}
	err = row.Scan(&record.Id, &record.IsExpense, 
		&record.Amount, &record.TypeName, 
		&record.Description, &record.Date)
	if err != nil {
		log.Fatal("for rows error", err)
	}
	return record
}

func Del(k string) {
	log.Println("Del transaction")
	delete := `DELETE FROM transactions WHERE id=$1;`
	id, err := strconv.Atoi(k)
	if err != nil {
		fmt.Printf("convert id error:%v\n", err)
		log.Fatal("convert id error:", err)
	}
	result, err := db.DB.Exec(delete, id)
	if err != err {
		fmt.Printf("delete row's id:%v, error:%v\n", id, err)
		log.Fatal("delete row's id:", id, " fail, err:", err)
	}
	rf, err := result.RowsAffected()
	log.Printf("RowsAffected:%v, error:%v\n", rf, err)
}

func Edit(r Record) {
	fmt.Printf("record: %#v\n", r)
	sqlStatement := `
	UPDATE transactions
	SET is_expense = $2 amount = $3 type_name = $4 description = $5 date = $6
	WHERE id = $1;`
	result, err1 := db.DB.Exec(sqlStatement, r.Id, r.IsExpense, r.Amount, r.TypeName, r.Description, r.Date)
	if err1 != nil {
		fmt.Printf("edit row's id:%v, result: %#v, error:%#v\n", r.Id, result, err1)
		log.Fatal("edit row's id:", r.Id, " error:", err1)
	}
	rf, err := result.RowsAffected()
	fmt.Printf("RowsAffected:%v, error:%v\n", rf, err)
}