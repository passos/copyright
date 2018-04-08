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
func (acc *Account) AddAccount() (val bool, err error) {
	res, err := db.Exec("insert into account(email,username,identity_id) values(?,?,?)", acc.Email, acc.Username, acc.IdentityID)
	if err != nil {
		val = false
		fmt.Println("called AddAccount() err")
		return false, err
	}
	id, err := res.LastInsertId()
	acc.ID = int(id)

	return true, err
}

//通用查询，返回map嵌套map
func query(sql string) map[int]map[string]string {
	fmt.Println("query is called")
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("query data err", err)
		return nil
	}
	//得到列名数组
	cols, err := rows.Columns()
	//获取列的个数
	colCount := len(cols)
	values := make([]string, colCount)
	oneRows := make([]interface{}, colCount)
	for k, _ := range values {
		oneRows[k] = &values[k] //将查询结果的返回地址绑定，这样才能变参获取数据
	}
	//存储最终结果
	results := make(map[int]map[string]string)
	idx := 0
	for rows.Next() {
		rows.Scan(oneRows...)
		rowmap := make(map[string]string)
		for k, v := range values {
			rowmap[cols[k]] = v

		}
		results[idx] = rowmap
		idx++
		//fmt.Println(values)
	}
	//fmt.Println("---------------------------------------")
	//fmt.Println(results)
	return results

}
