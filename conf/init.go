package conf

import (
	"fmt"
	"go-memo/cache"
	"go-memo/model"
	"gopkg.in/ini.v1"
)

// @Author KHighness
// @Update 2022-02-13

var (
	AppMode      string
	HttpPort     string
)

func Init()  {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		fmt.Println("Load config file error, please check file path")
		panic(err)
	}
	LoadServerConfig(file)
	model.MySQL(file)
	cache.Redis(file)
}

func LoadServerConfig(file *ini.File) {
	server := file.Section("server")
	AppMode = server.Key("AppMode").String()
	HttpPort = server.Key("HootPort").String()
}
