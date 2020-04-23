package tools

import (
	"bufio"
	"encoding/json"
	"os"
)

type (
	Config struct {
		AppName string `json:"app_name"`
		AppHost string `json:"app_host"`
		AppPort string `json:"app_port"`
		DatabaseConfig DatabaseConfig `json:"database_config"`
		RedisConfig RedisConfi `json:"redis_config"`
	}
	DatabaseConfig struct {
		Driver string `json:"driver"`
		User string `json:"user"`
		Password string `json:"password"`
		Host string `json:"host"`
		Port string `json:"port"`
		Database string `json:"database"`
		Charset string `json:"charset"`
		Sqlstr bool `json:"sqlstr"`
	}
	RedisConfi struct {
		Addr string `json:"addr"`
		Port string `json:"port"`
		Password string `json:"password"`
		Db int `json:"db"`
	}
)
var Cfg *Config = nil
//读取文件
func ReadConfig(path string)(*Config,error)  {
	file,err := os.Open(path)
	if err!=nil{
		panic(err)
	}
	//关闭文件
	defer file.Close()
	//读取文件
	readers := bufio.NewReader(file)
	//转义json
	decode :=json.NewDecoder(readers)
	//赋值
	if err = decode.Decode(&Cfg);err!=nil{
		return nil,err
	}
	return Cfg,nil
}