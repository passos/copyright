package main

import (
	"database/sql"
	"fmt"

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
func (this *Account) AddAccount() (val bool, err error) {
	_, err = db.Exec("insert into account(email,username,identity_id) values(?,?,?)", this.Email, this.Username, this.IdentityID)
	if err != nil {
		val = false
		fmt.Println("called AddAccount() err")
		return false, err
	}
	return true, nil
}

func query(sql string) {
	//rows, err := db.Query("SELECT * FROM account")
	//
	//// Get column names
	//columns, err := rows.Columns()
	//values := make([]sql.RawBytes, len(columns))
}
