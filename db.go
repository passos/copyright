package main

import (
	"database/sql"
	"fmt"
	"strconv"

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
	fmt.Println(results, "idx===", idx)
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
func getMatureAution() error {
	//查询account_content表中status为1且时间到的记录
	aution := &Aution{}
	//sql := "select content_hash,account_id,percent,sell_percent,sell_price from account_content where status ='1' and date_add(ts,interval 1 day) < now() and percent > 0"
	sql := "select content_hash,account_id,percent,sell_percent,sell_price from account_content where status ='1' and  percent > 0"

	m, _, err := query(sql)
	if err != nil {
		fmt.Println("db err,query account_content", err)
		return err
	}
	for _, v := range m {
		aution.AccountID, _ = strconv.Atoi(v["account_id"])
		aution.Content_hash = v["content_hash"]
		aution.Percent, _ = strconv.Atoi(v["percent"])
		aution.SellPercent, _ = strconv.Atoi(v["sell_percent"])
		aution.SellPrice, _ = strconv.Atoi(v["sell_price"])
		//调用单个发起拍卖方的竞价处理 -成交的可能必须是卖多少买多少
		aution.DealOneAution()
	}
	return err
}

//处理交易
func (aut *Aution) DealOneAution() error {
	//需要查询用户拍卖求购信息
	sql := fmt.Sprintf("select * from aution where  content_hash = '%s' order by price desc limit 1", aut.Content_hash)
	m, _, err := query(sql)
	if err != nil {
		fmt.Println("db err,query account_content", err)
		return err
	}
	leek := &Aution{}
	trades := make(map[int]*Trade)
	//	left_percent := aut.SellPercent

	trades[aut.AccountID] = &Trade{}
	trades[aut.AccountID].Percent = aut.SellPercent
	//aut.Percent = left_percent
	for _, v := range m {
		leek.AccountID, _ = strconv.Atoi(v["account_id"])
		leek.Percent, _ = strconv.Atoi(v["percent"])
		leek.Price, _ = strconv.Atoi(v["price"])
		//判断是否符合交易
		if leek.Price == aut.SellPrice {
			leek.Percent += aut.SellPercent
			aut.Percent -= aut.SellPercent
			sql = fmt.Sprintf("update account_content set percent=%d,sell_price=0,sell_percent=0,status='0' where content_hash='%s' and account_id=%d", aut.Percent, aut.Content_hash, aut.AccountID)
			if _, err = Create(sql); err != nil {
				fmt.Println("update account_content err", err)
				return err
			}
			sql = fmt.Sprintf("insert into account_content(account_id,content_id,content_hash,percent) values(%d,%d,'%s',%d)", leek.AccountID, aut.ContentID, aut.Content_hash, leek.Percent)
			if _, err = Create(sql); err != nil {
				fmt.Println("update account_content err", err)
				return err
			}
			break
		}
		//		if left_percent > 0 && leek.Percent > 0 {
		//			trades[leek.AccountID] = &Trade{}
		//			if left_percent < leek.Percent {
		//				trades[aut.AccountID].Amount += left_percent * leek.Price
		//				left_percent -= left_percent
		//				trades[aut.AccountID].Percent -= left_percent
		//				trades[leek.AccountID].Amount -= left_percent * leek.Price
		//				trades[leek.AccountID].Percent += left_percent
		//			} else {
		//				trades[aut.AccountID].Amount += leek.Percent * leek.Price
		//				left_percent -= left_percent
		//				trades[aut.AccountID].Percent -= leek.Percent
		//				trades[leek.AccountID].Amount -= leek.Percent * leek.Price
		//				trades[leek.AccountID].Percent += leek.Percent
		//			}
		//		}
		//		if left_percent == 0 {
		//			break
		//		}

	}
	//	aut.Percent = left_percent
	//根据成交结果更新数据库
	//	for k, v := range trades {
	//		if k == aut.AccountID {
	//			sql = fmt.Sprintf("update account_content set percent=%d,sell_price=0,sell_percent=0 where content_hash='%s' and account_id=%d", v.Percent, aut.Content_hash, aut.AccountID)
	//			if _, err = Create(sql); err != nil {
	//				fmt.Println("update account_content err", err)
	//				return err
	//			}
	//		} else {
	//			sql = fmt.Sprintf("insert into account_content(account_id,content_id,content_hash,percent) values(%d,%d,'%s',%d)", k, aut.ContentID, aut.Content_hash, v.Percent)
	//			if _, err = Create(sql); err != nil {
	//				fmt.Println("insert into account_content err", err)
	//				return err
	//			}
	//		}

	//	}
	return err
}
