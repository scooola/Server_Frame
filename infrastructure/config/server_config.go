//Auther: scola
//Date: 2021/08/08 20:15
//Description:
//Σ(っ °Д °;)っ

package config

import (
	"Server_Frame/framework/errors"
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
)

type ServerConfig struct {
	Redis RedisConfig `json:redis`
}

type RedisConfig struct {
	Host     string `json:host`
	Port     int    `json:port`
	DB       int    `json:db`
	PassWord string `json:password`
}

func GetJsonConf() ServerConfig {
	curPath, _ := os.Getwd()
	sysType := runtime.GOOS
	var jsonPath string
	if sysType == "linux" {
		jsonPath = curPath + "/server_config.json"
	} else {
		jsonPath = curPath + "\\server_config.json"
	}

	jsonFile, err := os.Open(jsonPath)

	if err != nil {
		errors.RaiseErr(err.Error())
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var conf ServerConfig

	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		errors.RaiseErr(err.Error())
	}

	return conf
}
