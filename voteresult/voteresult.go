//处理投票最多的
package main

import (
	"copyright/abi"
	"copyright/configs"
	"copyright/dbs"
	"fmt"
	"strconv"
	"strings"
	"time"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var config *configs.ServerConfig

const MAX_ONE_TRANSFER = 100 //批量转账一次最多转的个数

func main() {
	config = configs.GetConfig()
	if config == nil {
		return
	}
	fmt.Println("get config ", config)
	fmt.Println("rpc==", config.Eth)
	dbs.DBConn = dbs.InitDB(config)
	//每天启动时计算之前一天投票的结果
	GetDayList()

	fmt.Println(time.Now().Weekday().String())
	//每周启动
	if time.Now().Weekday() == time.Monday {
		go GetWeekList()
	}
}

//获得日排行榜
func GetDayList() error {
	data, _, err := dbs.DBQuery(`select a.content_hash,count(*) cnt,c.address 
from vote a,account_content b,account c 
where date_format(vote_time,'%Y-%m-%d') = date_format(date_sub(curdate(),INTERVAL 1 DAY),'%Y-%m-%d')
  and a.content_hash = b.content_hash 
  and b.account_id =c.account_id 
group by a.content_hash ,c.address 
limit 10`)
	if err != nil {
		fmt.Println("query data err")
		return err
	}
	for _, v := range data {
		//更新数据库-记录用户资产投票数量
		voting, _ := strconv.Atoi(v["cnt"])
		_, err = dbs.DBConn.Exec("update account_content set voting = ? where content_hash = ?", voting, v["content_hash"])
		if err != nil {
			fmt.Println("GetWeekList update account_content err:", err)
			break
		}
		//针对前10名发去贺电 -- 包括图片拥有者和图片投票者
		//给图片拥有者发出贺电
		transfer20("yekai", config.Eth.MgrAddress, v["address"], int64(voting/10000+100))
		//查询每个图片的投票者，送出贺电
		//select b.address from vote a,account b where a.content_hash = 'xx' and a.account_id=b.account_id;
		rows, err := dbs.DBConn.Query("select b.address from vote a,account b where a.content_hash = ? and a.account_id=b.account_id", v["content_hash"])
		if err != nil {
			fmt.Println("GetWeekList query voter err:", err)
			break
		}
		var addr string
		addrs := make([]string, 200)
		for rows.Next() {
			err = rows.Scan(&addr)
			if err != nil {
				break
			}
			addrs = append(addrs, addr)
			//具体转账算法需要仔细核算，本次只是举个栗子
			//transfer20("yekai", config.Eth.MgrAddress, addr, int64(voting/100000+100))

		}
		//multiTransfer(frompass, fromaddr string, toaddrs []string, amount int64)
		multiTransfer("yekai", config.Eth.MgrAddress, addrs, int64(voting/100000+100))
	}

	return err
}

//获得周排行榜
func GetWeekList() error {
	data, _, err := dbs.DBQuery(`select a.content_hash,count(*) cnt,c.address 
from vote a,account_content b,account c 
where date_format(vote_time,'%Y-%m-%d') between date_sub(curdate(),INTERVAL WEEKDAY(curdate()) + 7 DAY) and date_sub(curdate(),INTERVAL WEEKDAY(curdate()) + 1 DAY) 
  and a.content_hash = b.content_hash 
  and b.account_id =c.account_id 
group by a.content_hash ,c.address 
limit 10`)
	if err != nil {
		fmt.Println("query data err")
		return err
	}
	for _, v := range data {
		//更新数据库-记录用户资产投票数量
		voting, _ := strconv.Atoi(v["cnt"])
		_, err = dbs.DBConn.Exec("update account_content set voting = ? where content_hash = ?", voting, v["content_hash"])
		if err != nil {
			fmt.Println("GetWeekList update account_content err:", err)
			break
		}
		//针对前10名发去贺电 -- 包括图片拥有者和图片投票者
		//给图片拥有者发出贺电
		transfer20("yekai", config.Eth.MgrAddress, v["address"], int64(voting/10000+100))
		//查询每个图片的投票者，送出贺电
		//select b.address from vote a,account b where a.content_hash = 'xx' and a.account_id=b.account_id;
		rows, err := dbs.DBConn.Query("select b.address from vote a,account b where a.content_hash = ? and a.account_id=b.account_id", v["content_hash"])
		if err != nil {
			fmt.Println("GetWeekList query voter err:", err)
			break
		}
		var addr string
		addrs := make([]string, 200)
		for rows.Next() {
			err = rows.Scan(&addr)
			if err != nil {
				break
			}
			addrs = append(addrs, addr)

		}
		//具体转账算法需要仔细核算，本次只是举个栗子
		multiTransfer("yekai", config.Eth.MgrAddress, addrs, int64(voting/100000+100))
	}

	return err
}

//erc20转账
func transfer20(frompass, fromaddr, toaddr string, amount int64) error {
	//ipc文件地址
	conn, err := ethclient.Dial(config.Eth.Ipcfile)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		return err
	}
	//通过合约地址获得入口
	token, err := abi.NewPxcoin(common.HexToAddress(config.Eth.Contract20), conn)
	if err != nil {
		fmt.Println("Failed to instantiate a Token contract: ", err)
		return err
	}
	total, err := token.TotalSupply(nil)
	if err != nil {
		fmt.Println("Failed to retrieve token name: ", err)
		return err
	}
	fmt.Println("Total:", total)
	//通过地址获得文件信息
	keyname, err := abi.GetFileName(string([]rune(fromaddr)[2:]), config.Eth.Keydir)
	if err != nil || keyname == "" {
		fmt.Println("Pic721Token:Failed to get key file Name: ", err, config.Eth.Keydir, fromaddr)
		return err
	}
	key, err := abi.ReadKeyFile(config.Eth.Keydir + keyname)
	if err != nil || keyname == "" {
		fmt.Printf("Pic721Token:Failed to read key file: %v\n", err)
		return err
	}
	// 设置转账者
	auth, err := bind.NewTransactor(strings.NewReader(key), frompass)
	if err != nil {
		fmt.Println("Failed to create authorized transactor: ", err)
		return err
	}
	//先查余额
	balance, err := token.BalanceOf(nil, common.HexToAddress(toaddr))

	if err != nil {
		fmt.Println("Failed to get balance: ", err)
		return err
	}
	fmt.Println("balance of toaddr:", balance)
	//构造转账金额
	num := big.NewInt(amount)

	trans1, err1 := token.Transfer(auth, common.HexToAddress(toaddr), num)

	if err1 != nil {
		fmt.Println("Failed to transfer: ", err)
		return err
	}
	fmt.Println("trans1:", trans1)
	balance, err = token.BalanceOf(nil, common.HexToAddress(toaddr))
	if err != nil {
		fmt.Println("Failed to get balance of key1: ", err)
		return err
	}

	fmt.Println("balance of toaddr:", balance)
	return err
}

//批量转账
func multiTransfer(frompass, fromaddr string, toaddrs []string, amount int64) error {
	//ipc文件地址
	conn, err := ethclient.Dial(config.Eth.Ipcfile)
	if err != nil {
		fmt.Println("multiTransfer:Failed to connect to the Ethereum client:", err)
		return err
	}
	//通过合约地址获得入口
	token, err := abi.NewPxcoin(common.HexToAddress(config.Eth.Contract20), conn)
	if err != nil {
		fmt.Println("multiTransfer:Failed to instantiate a Token contract: ", err)
		return err
	}
	//通过地址获得文件信息
	keyname, err := abi.GetFileName(string([]rune(fromaddr)[2:]), config.Eth.Keydir)
	if err != nil || keyname == "" {
		fmt.Println("multiTransfer:Failed to get key file Name: ", err, config.Eth.Keydir, fromaddr)
		return err
	}
	key, err := abi.ReadKeyFile(config.Eth.Keydir + keyname)
	if err != nil || keyname == "" {
		fmt.Printf("multiTransfer:Failed to read key file: %v\n", err)
		return err
	}
	// 设置转账者
	auth, err := bind.NewTransactor(strings.NewReader(key), frompass)
	if err != nil {
		fmt.Println("Failed to create authorized transactor: ", err)
		return err
	}
	index := 0
	addr := make([]common.Address, MAX_ONE_TRANSFER)
	for _, v := range toaddrs {
		addr[index] = common.HexToAddress(v)
		index++
		if index == MAX_ONE_TRANSFER {
			//批量提交一次
			token.TranferAll(auth, addr, big.NewInt(amount))
			index = 0
		}
	}
	//防止侧漏
	if index > 0 {
		token.TranferAll(auth, addr, big.NewInt(amount))
	}
	return err
}
