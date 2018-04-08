package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@Bailian123@localhost/copyright")
	if err != nil {
		panic(err.Error())
	}
	return db, err
}
