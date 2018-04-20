package main

import (
	"fmt"

	"crypto/md5"

	"github.com/ethereum/go-ethereum/rpc"
)

type Block struct {
	Number string
}

func GetAccAddress(pass, connstr string) (string, error) {
	// Connect the client
	client, err := rpc.Dial(connstr)
	//client, err := rpc.Dial("http://192.168.137.131:8000")
	if err != nil {
		//log.Fatalf("could not create ipc client: %v", err)
		fmt.Println("connect to localrpc", err)
		return "", err
	}

	//var lastBlock Block
	var account string
	//err = client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
	err = client.Call(&account, "personal_newAccount", pass)

	if err != nil {
		fmt.Println("can't get latest block:", err)
		return "", err
	}

	// Print events from the subscription as they arrive.
	fmt.Printf("account: %s\n", account)
	return account, nil
}

func GetMd5(pass string) string {
	ha := md5.New()
	ha.Write([]byte(pass))
	s := ha.Sum(nil)
	fmt.Printf("%x\n", s)
	return fmt.Sprintf("%x", s)
}
