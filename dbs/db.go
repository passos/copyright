package dbs

import (
	"copyright/configs"
	"database/sql"
	"fmt"
	_ "strconv"

	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

type Vote struct {
	AccountID    int    `json:"account_id"`
	Content_hash string `json:"content_hash"`
	Comment      string `json:"comment"`
}

func InitDB(config *configs.ServerConfig) *sql.DB {
	db, err := sql.Open(config.Db.Driver, config.Db.Url)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return db
}

//通用查询，返回map嵌套map
func DBQuery(sql string) ([]map[string]string, int, error) {
	fmt.Println("query is called", sql)
	rows, err := DBConn.Query(sql)
	if err != nil {
		fmt.Println("query data err", err)
		return nil, 0, err
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
	results := make([]map[string]string, 10)
	idx := 0
	for rows.Next() {
		rows.Scan(oneRows...)
		rowmap := make(map[string]string)
		for k, v := range values {
			rowmap[cols[k]] = v

		}
		results = append(results, rowmap)
		idx++
		//fmt.Println(values)
	}
	//fmt.Println("---------------------------------------")
	fmt.Println("query..idx===", idx)
	return results, idx, nil

}
func Create(sql string) (int64, error) {
	res, err := DBConn.Exec(sql)
	if err != nil {
		fmt.Println("exec sql err,", err, "sql is ", sql)
		return -1, err
	}
	return res.LastInsertId()
}

func (vot *Vote) AddVote() error {
	_, err := DBConn.Exec("insert into vote(account_id,content_hash,comment) values(?,?,?)", vot.AccountID, vot.Content_hash, vot.Comment)
	if err != nil {
		fmt.Println("add a vote err:", err)
		return err
	}
	return nil
}
