package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB(config *ServerConfig) *sql.DB {
	db, err := sql.Open(config.Db.Driver, config.Db.Url)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return db
}

func query(sql string) {
	//rows, err := db.Query("SELECT * FROM account")
	//
	//// Get column names
	//columns, err := rows.Columns()
	//values := make([]sql.RawBytes, len(columns))
}
