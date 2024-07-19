package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

const transaction string = `
CREATE TABLE IF NOT EXISTS transactions (
	id INT AUTO_INCREMENT,
	is_expense BOOLEAN NOT NULL,
	amount FLOAT NOT NULL,
	type_name TEXT NOT NULL,
	description TEXT,
	date TEXT NOT NULL,
	PRIMARY KEY (id)
)
`

func InitDB() {
	var err error
	DB, err = sql.Open("ramsql", "accounting")
	if err != nil {
		log.Fatal("error", err)
	}

	initTable(transaction)
}

func initTable(sql string) {
	var err error
	_, err = DB.Exec(sql)
	if err != nil {
		log.Fatal("error:", err)
	}
}