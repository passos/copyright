package main

import (
	"copyright/abi"
	"fmt"
	"math/big"
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
	token, err := abi.NewErc20(common.HexToAddress(config.Eth.Contract20), conn)
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

	keyname, err := abi.GetFileName(config.Eth.MgrAddress, config.Eth.Keydir)
	if err != nil || keyname == "" {
		fmt.Println("InitAccToken :Failed to get key file Name: ", err, config.Eth.MgrAddress, keyname)
		return err
	}
	key, err := abi.ReadKeyFile(config.Eth.Keydir + keyname)
	if err != nil || keyname == "" {
		fmt.Printf("InitAccToken:Failed to read key file: %v\n", err)
		return err
	}

	//	data, err := ioutil.ReadFile(config.Eth.Key)
	//	if err != nil {
	//		fmt.Printf("Failed to read key files: %v,file[%s]\n", err, config.Eth.Key)
	//		return err
	//	}
	//	key := string(data)

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
	token, err := abi.NewErc20(common.HexToAddress(config.Eth.Contract20), conn)
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

//上传图片资产获得图片token
func Pic721Token(pichash, address, contractaddress, pass string) error {
	conn, err := ethclient.Dial(config.Eth.Ipcfile)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v\n", err)
		return err
	}
	// Instantiate the contract and display its name
	//合约地址
	token, err := abi.NewErc721(common.HexToAddress(contractaddress), conn)

	if err != nil {
		fmt.Printf("Failed to instantiate a Token contract: %v\n", err)
		return err
	}
	total, err := token.TotalSupply(nil)
	if err != nil {
		fmt.Printf("Failed to retrieve token name: %v\n", err)
		return err
	}
	fmt.Println("Total:", total)
	//需要根据address读取
	keyname, err := abi.GetFileName(address, config.Eth.Keydir)
	if err != nil || keyname == "" {
		fmt.Println("Pic721Token:Failed to get key file Name: ", err, config.Eth.Keydir, address)
		return err
	}
	key, err := abi.ReadKeyFile(config.Eth.Keydir + keyname)
	if err != nil || keyname == "" {
		fmt.Printf("Pic721Token:Failed to read key file: %v\n", err)
		return err
	}
	// Create an authorized transactor and spend 1 unicorn
	//需要用户密码
	auth, err := bind.NewTransactor(strings.NewReader(key), pass)
	if err != nil {
		fmt.Printf("Failed to create authorized transactor: %v\n", err)
		return err
	}

	//auth1, err := bind.NewTransactor(strings.NewReader(key1), "123456")
	//num := big.NewInt(tokenid)

	trans1, err := token.NewToken(auth, pichash, pichash)

	if err != nil {
		fmt.Printf("Failed to transfer: %v\n", err)
		return err
	}
	fmt.Println("trans1:", trans1)

	asset_num, err := token.TotalSupply(nil)
	if err != nil {
		fmt.Printf("Failed to get number of asset: %v\n", err)
		return err
	}
	// need to wait some time for the transaction commit
	fmt.Println("Total after:", asset_num)
	return nil
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
	token, err := abi.NewErc20(common.HexToAddress(config.Eth.Contract20), conn)
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

//erc721交易执行
func transfer721(frompass, fromaddr, toaddr string, tokenid int64) error {
	//connect by ipcfile
	conn, err := ethclient.Dial(config.Eth.Ipcfile)
	if err != nil {
		fmt.Println("transfer721:Failed to connect to the Ethereum client: ", err)
		return err
	}
	// get in by contract addr
	token, err := abi.NewErc721(common.HexToAddress(config.Eth.Contract721), conn)

	if err != nil {
		fmt.Println("transfer721:Failed to instantiate a Token contract: ", err)
		return err
	}

	//获得资产id
	asset_num, err := token.TotalSupply(nil)
	if err != nil {
		fmt.Println("Failed to retrieve token name: ", err)
		return err
	}
	fmt.Println("transfer721----asset_num:", asset_num)
	//通过地址获得文件信息
	keyname, err := abi.GetFileName(fromaddr, config.Eth.Keydir)
	if err != nil || keyname == "" {
		fmt.Println("transfer721:Failed to get key file Name: ", err, config.Eth.Keydir, fromaddr)
		return err
	}
	key, err := abi.ReadKeyFile(config.Eth.Keydir + keyname)
	if err != nil || keyname == "" {
		fmt.Printf("transfer721:Failed to read key file: %v\n", err)
		return err
	}
	// Create an authorized transactor and spend 1 unicorn
	auth, err := bind.NewTransactor(strings.NewReader(key), frompass)
	if err != nil {
		fmt.Println("transfer721:Failed to create authorized transactor: ", err)
		return err
	}
	//tokenId := big.NewInt(3)

	trans1, err := token.Transfer(auth, common.HexToAddress("0x5b2274dd73ffd4fdf154b521b84b1834e18c2fae"), big.NewInt(tokenid))

	if err != nil {
		fmt.Println("transfer721:Failed to transfer token:", err)
		return err
	}
	// need to wait some time for the transaction commit
	fmt.Println("trans1:", trans1)

	//	total1, err := token.TokensOfOwner(nil, common.HexToAddress("0x5b2274dd73ffd4fdf154b521b84b1834e18c2fae"))
	//	if err != nil {
	//		fmt.Println("transfer721:Failed to token of owner: ", err)
	//		return err
	//	}
	//	fmt.Println("total after transfer:", total1)
	return err

}

func AssetSplit721(fromaddr, frompass string, tokenid, weight int64) (int, error) {
	//connect by ipcfile
	conn, err := ethclient.Dial(config.Eth.Ipcfile)
	if err != nil {
		fmt.Println("transfer721:Failed to connect to the Ethereum client: ", err)
		return 0, err
	}
	// get in by contract addr
	token, err := abi.NewErc721(common.HexToAddress(config.Eth.Contract721), conn)

	if err != nil {
		fmt.Println("AssetSplit721:Failed to instantiate a Token contract: ", err)
		return 0, err
	}
	//通过地址获得文件信息
	keyname, err := abi.GetFileName(string([]rune(fromaddr)[2:]), config.Eth.Keydir)
	if err != nil || keyname == "" {
		fmt.Println("AssetSplit721:Failed to get key file Name: ", err, config.Eth.Keydir, fromaddr)
		return 0, err
	}
	key, err := abi.ReadKeyFile(config.Eth.Keydir + keyname)
	if err != nil || keyname == "" {
		fmt.Printf("AssetSplit721:Failed to read key file: %v\n", err)
		return 0, err
	}
	// Create an authorized transactor and spend 1 unicorn
	auth, err := bind.NewTransactor(strings.NewReader(key), frompass)
	if err != nil {
		fmt.Println("AssetSplit721:Failed to create authorized transactor: ", err)
		return 0, err
	}
	_, err = token.Split(auth, big.NewInt(tokenid), big.NewInt(weight))
	if err != nil {
		fmt.Println("AssetSplit721:Failed to Split transactor: ", err)
		return 0, err
	}
	return 0, err
}
