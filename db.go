package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

var db *sql.DB

type (
	Account struct {
		ID         int    `json:"id"`
		Email      string `json:"email"`
		Username   string `json:"username"`
		IdentityID string `json:"identity_id"`
	}
)

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

func dbQuery(query string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(query, args...)

	if err != nil {
		log.Error(err)
	}

	// Get column names
	columns, err := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
}
