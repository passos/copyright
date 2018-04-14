package abi

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GetFileName(address, dirname string) (string, error) {
	data, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println("read dir err", err)
		return "", err
	}
	for _, v := range data {
		if strings.Index(v.Name(), address) > 0 {
			//代表找到文件
			return v.Name(), nil
		}
	}

	return "", nil
}

func ReadKeyFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("faile to read file:", err)
		return "", err
	}
	return string(data), nil
}

//func main() {
//	xx := getFileName("721", "./")
//	fmt.Println(xx)
//}
