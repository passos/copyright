package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type ServerConfig struct {
	Common *CommonConfig
	Db     *DbConfig
	Eth    *EthConfig
}

type CommonConfig struct {
	Port      int
	LogFormat string
}

type DbConfig struct {
	Driver string
	Url    string
}

type EthConfig struct {
	Rpc         string
	Contract20  string
	Contract721 string
	Ipcfile     string
	Keydir      string
	MgrAddress  string
}

func decodeStr(str string) string {
	if strings.Trim(str, " ") == "" {
		return ""
	}

	for i := 0; i < 3; i++ {
		bytes, _ := base64.StdEncoding.DecodeString(str[1:])
		str = string(bytes)
	}
	return str
}

func encodeStr(str string) string {
	for i := 0; i < 3; i++ {
		str = string(str[7*i%len(str)]) + base64.StdEncoding.EncodeToString([]byte(str))
	}
	return str
}

func usage() {
	fmt.Printf("Usage: %s -c config_file [-v] [-h]\n", os.Args[0])
}

func getConfig() (config *ServerConfig) {
	var configFile = flag.String("c", "", "Config file")
	var encode = flag.String("e", "", "encode")
	var decode = flag.String("d", "", "decode")

	var ver = flag.Bool("v", false, "version")
	var help = flag.Bool("h", false, "Help")

	flag.Usage = usage
	flag.Parse()

	if *help {
		usage()
		return nil
	}

	if *ver {
		fmt.Println("Version: ", Version)
		fmt.Println("Commit: ", Commit)
		fmt.Println("BuildTime: ", BuildTime)
		return nil
	}

	if *encode != "" {
		fmt.Println(encodeStr(*encode))
		return nil
	}

	if *decode != "" {
		fmt.Println(decodeStr(*decode))
		return nil
	}

	// get server config
	if *configFile != "" {
		config = &ServerConfig{}
		if _, err := toml.DecodeFile(*configFile, &config); err != nil {
			panic(err)
		}
	} else {
		config = &ServerConfig{
			Common: &CommonConfig{
				Port:      9080,
				LogFormat: `${time_rfc3339} [${prefix}] [${level}]`,
				//LogFormat: `${time_rfc3339} ${short_file}:${line} [${prefix}] [${level}]`,
			},
		}
	}

	return config
}
