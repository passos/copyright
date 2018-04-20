//定时处理任务-周期性处理拍卖结果
package main

import (
	"copyright/configs"
	"fmt"
	"math/big"
	"time"

	"copyright/abi"
	"copyright/dbs"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Aution struct {
	AccountID    int    `json:"account_id"`
	ContentID    int    `json:"content_id"`
	Content_hash string `json:"content_hash"`
	Percent      int    `json:"percent"`
	Price        int    `json:"price"`
	SellPrice    int    `json:"sell_price"`
	SellPercent  int    `json:"sell_percent"`
	Status       int    `json:"status"`
	Address      string `json:"address"`
	TokenID      int64  `json:"tokenid"`
}

var config *configs.ServerConfig

func main() {
	config = configs.GetConfig()
	if config == nil {
		return
	}
	fmt.Println("get config ", config)
	fmt.Println("rpc==", config.Eth)
	dbs.DBConn = dbs.InitDB(config)
	//周期性启动去检查任务
	tk := time.NewTicker(time.Second * 1800) //半小时的tk
	go getMatureAution()
	for {
		<-tk.C //阻塞等待管道
		go getMatureAution()
	}
}
func getTokenID(accaddr, pixhash string) (int64, error) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Println("getTokenID Dial err:", err)
		return 0, err
	}
	token, err := abi.NewPixasset(common.HexToAddress(config.Eth.Contract721), client)
	if err != nil {
		fmt.Println("getTokenID NewPixasset err:", err)
		return 0, err
	}
	data, err := token.FindTokenId(nil, pixhash, common.HexToAddress(accaddr))
	if err != nil {
		fmt.Println("getTokenID FindTokenId err:", err)
		return 0, err
	}
	return data.Int64(), err

}
func TransferFrom(from, to string, val int64) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Println("TransferFrom Dial err:", err)
		return 0, err
	}
	token, err := abi.NewPxcoin(common.HexToAddress(config.Eth.Contract20), client)
	if err != nil {
		fmt.Println("TransferFrom NewPxcoin err:", err)
		return 0, err
	}
	num := big.NewInt(val)
	_, err = token.TransferFrom(nil, common.HexToAddress(from), common.HexToAddress(to), num)
	if err != nil {
		fmt.Println("TransferFrom   err:", err)
		return 0, err
	}
}
func getMatureAution() error {
	//查询account_content表中status为1且时间到的记录
	aution := &Aution{}
	//sql := "select content_hash,account_id,percent,sell_percent,sell_price from account_content where status ='1' and date_add(ts,interval 1 day) < now() and percent > 0"
	sql := "select content_hash,a.account_id,percent,sell_percent,sell_price,content_id,b.identity_id pass,b.address,a.tokenid from account_content a,account b where status ='1' and  percent > 0 and a.account_id=b.account_id"

	m, _, err := dbs.DBQuery(sql)
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
		aution.ContentID, _ = strconv.Atoi(v["content_id"])

		aution.Address = v["address"]
		aution.TokenID, _ = getTokenID(aution.Address, aution.Content_hash)
		//调用单个发起拍卖方的竞价处理 -成交的可能必须是卖多少买多少
		aution.DealOneAution()
	}
	return err
}

//处理交易
func (aut *Aution) DealOneAution() error {
	//需要查询用户拍卖求购信息
	fmt.Println("DealOneAution run ...hash===", aut.Content_hash)
	sql := fmt.Sprintf("select a.content_hash,a.account_id,a.percent,a.price,b.address,b.identity_id pass from aution a,account b  where a.account_id=b.account_id and  content_hash = '%s' order by price desc limit 1", aut.Content_hash)
	m, _, err := dbs.DBQuery(sql)
	if err != nil {
		fmt.Println("db err,query account_content", err)
		return err
	}
	leek := &Aution{}
	//aut.Percent = left_percent
	for _, v := range m {
		leek.AccountID, _ = strconv.Atoi(v["account_id"])
		//leek.Percent, _ = strconv.Atoi(v["percent"])
		leek.Price, _ = strconv.Atoi(v["price"])
		leek.Address = v["address"]
		//判断是否符合交易
		if leek.Price == aut.SellPrice {
			//			if aut.SellPercent < aut.Percent {
			//				//需要资产分割
			//				newTokenID := AssetSplit721(aut.Address, aut.Pass, aut.TokenID, aut.SellPercent)
			//			}
			leek.Percent += aut.SellPercent
			aut.Percent -= aut.SellPercent
			sql = fmt.Sprintf("update account_content set percent=%d,sell_price=0,sell_percent=0,status='0' where content_hash='%s' and account_id=%d", aut.Percent, aut.Content_hash, aut.AccountID)
			if _, err = dbs.Create(sql); err != nil {
				fmt.Println("update account_content err", err)
				return err
			}
			fmt.Println("leek is :", leek)
			sql = fmt.Sprintf("insert into account_content(account_id,content_id,content_hash,percent) values(%d,%d,'%s',%d)", leek.AccountID, aut.ContentID, aut.Content_hash, leek.Percent)
			if _, err = dbs.Create(sql); err != nil {
				fmt.Println("update account_content err", err)
				return err
			}
			//删除竞拍信息
			sql = fmt.Sprintf("delete from aution where content_hash ='%s'", aut.Content_hash)
			if _, err = dbs.Create(sql); err != nil {
				fmt.Println("failed to delete aution:", err)
				return err
			}
			break
			//调用智能合约完成交易闭环
			//			go func() {
			//				//涉及到pixc资产转移 -- 需要发起方的密码 -- 使用approve
			//				//transfer20(leek.Pass, leek.Address, aut.Address, int64(aut.SellPercent*aut.Price))
			//				TransferFrom(leek.Address, aut.Address, int64(aut.SellPercent*aut.Price))
			//				//涉及到图片资产转移
			//				if aut.SellPercent < aut.Percent {
			//					//涉及到资产分割之后再转移
			//				} else {
			//					//可以直接资产转移
			//					transfer721(aut.Pass, aut.Address, leek.Address, aut.TokenID)
			//				}

			//				//涉及到交易收取手续费，收取卖家
			//				transfer20(leek.Pass, aut.Address, config.Eth.MgrAddress, int64(aut.SellPercent*aut.Price*2/100))

			//			}() // --暂不调用

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
