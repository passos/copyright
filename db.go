package main

import (
	"copyright/configs"
	"database/sql"
	"fmt"
	_ "strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

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
func (acc *Account) AddAccount() (val bool, err error) {
	res, err := db.Exec("insert into account(email,username,identity_id,address) values(?,?,?,?)", acc.Email, acc.Username, acc.IdentityID, acc.Address)
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
func query(sql string) (map[int]map[string]string, int, error) {
	fmt.Println("query is called", sql)
	rows, err := db.Query(sql)
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
	fmt.Println("query..idx===", idx)
	return results, idx, nil

}
func Create(sql string) (int64, error) {
	res, err := db.Exec(sql)
	if err != nil {
		fmt.Println("exec sql err,", err, "sql is ", sql)
		return -1, err
	}
	return res.LastInsertId()
}
func (ctx *Content) AddContent() (int64, error) {
	res, err := db.Exec("insert into content(title,content,content_hash) values(?,?,?)",
		ctx.Title,
		ctx.Content,
		ctx.Content_hash,
	)
	if err != nil {
		fmt.Println("insert into content err", err)
		return -1, err
	}
	id, err := res.LastInsertId()
	ctx.ContentID = int(id)
	fmt.Println(" ctx.ContentID===", ctx.ContentID)
	return res.LastInsertId()
}
func (aut *Aution) AddAccountContent() (int64, error) {
	res, err := db.Exec("insert into account_content(account_id,content_id,content_hash,percent,sell_price,sell_percent) values(?,?,?,?,?,?)",
		aut.AccountID,
		aut.ContentID,
		aut.Content_hash,
		aut.Percent,
		aut.SellPrice,
		aut.SellPercent,
	)
	if err != nil {
		fmt.Println("insert into content err", err)
		return -1, err
	}
	return res.LastInsertId()
}
func (ctx *Content) getContent() error {
	row := db.QueryRow("select title,content from content where content_hash=?", ctx.Content_hash)
	if row == nil {
		fmt.Println("select content err")
		return nil
	}
	//ctx.Content = make([]byte, 19677)
	err := row.Scan(&ctx.Title, &ctx.Content)
	fmt.Println(err, ctx.Title, len(ctx.Content))
	return err
}

func (vot *Vote) AddVote() error {
	_, err := db.Exec("insert into vote(account_id,content_hash,comment) values(?,?,?)", vot.AccountID, vot.Content_hash, vot.Comment)
	if err != nil {
		fmt.Println("add a vote err:", err)
		return err
	}
	return nil
}
