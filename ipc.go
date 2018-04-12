package main

import (
	"fmt"
	"io/ioutil"
	//	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//const key = `{"address":"232db2a944a0c90b9855bbc2debad9e605aeda91","crypto":{"cipher":"aes-128-ctr","ciphertext":"aaeafdaedb5abd2bfe4bd2008fe96ff73c06777f5231941ea7f9876eb9aea431","cipherparams":{"iv":"90d1ae928166b7c110c14486e2a50dc0"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8ad5881e7e7383d1d7a84f2fcde2c8d6e12e79881964a6294269b108f4a125af"},"mac":"0a0482df86ac967dba9e6349002d5e12d89d3609bbac8dd676cdf8b2f54c8f83"},"id":"3b17c6a8-df65-452b-b849-170ca6606d11","version":3}`

func InitAccToken(address string) error {
	conn, err := ethclient.Dial(config.Eth.Ipcfile)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v\n", err)
		return err
	}
	//此处使用合约地址
	token, err := NewToken(common.HexToAddress(config.Eth.Contractaddr), conn)
	if err != nil {
		fmt.Printf("Failed to instantiate a Token contract: %v\n", err)
		return err
	}
	//准备发币
	total, err := token.TotalSupply(nil)
	if err != nil {
		fmt.Printf("Failed to retrieve token name: %v\n", err)
		return err
	}
	fmt.Println("Total:", total)
	data, err := ioutil.ReadFile(config.Eth.Key)
	if err != nil {
		fmt.Printf("Failed to read key files: %v,file[%s]\n", err, config.Eth.Key)
		return err
	}
	key := string(data)
	//设置管理者
	auth, err := bind.NewTransactor(strings.NewReader(key), "yekai")
	if err != nil {
		fmt.Printf("Failed to create authorized transactor: %v\n", err)
		return err
	}
	//设置收到的地址
	trans, err1 := token.InitialSupply(auth, common.HexToAddress(address))
	if err1 != nil {
		fmt.Printf("Failed to transfer: %v\n", err)
		return err
	}
	fmt.Println("trans:", trans)
	//查看收货地址的余额
	balance, err := token.BalanceOf(nil, common.HexToAddress(address))

	if err != nil {
		fmt.Printf("Failed to get balance: %v\n", err)
		return err
	}
	fmt.Println("balance:", balance)
	return nil
}

func GetBalanceOf(address string) (int64, error) {
	conn, err := ethclient.Dial(config.Eth.Ipcfile)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v\n", err)
		return 0, err
	}
	//此处使用合约地址
	token, err := NewToken(common.HexToAddress(config.Eth.Contractaddr), conn)
	if err != nil {
		fmt.Printf("Failed to instantiate a Token contract: %v\n", err)
		return 0, err
	}
	balance, err := token.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		fmt.Printf("Failed to get balance: %v\n", err)
		return 0, err
	}
	return balance.Int64(), err
}
