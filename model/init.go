package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
	"strings"
	"time"
)

// @Author KHighness
// @Update 2022-02-13

var (
	DB      *gorm.DB
	DbHost  string
	DbPort  string
	DbUser  string
	DbPass  string
	DbName  string
)

// 初始化MySQL
func MySQL(file *ini.File) {
	loadMysqlConfig(file)
	url := strings.Join([]string{DbUser, ":", DbPass, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	connectMySQL(url)
}

// 读取NySQL配置
func loadMysqlConfig(file *ini.File) {
	mysql := file.Section("mysql")
	DbHost = mysql.Key("DbHost").String()
	DbPort = mysql.Key("DbPort").String()
	DbUser = mysql.Key("DbUser").String()
	DbPass = mysql.Key("DbPass").String()
	DbName = mysql.Key("DbName").String()
}

// 连接到MySQL
func connectMySQL(url string)  {
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	if gin.Mode() == "release" {
		db.LogMode(false)
	} else {
		db.LogMode(true)
	}

	db.SingularTable(true)                    // 默认不加复数
	db.DB().SetMaxIdleConns(20)                   // 设置连接池的空闲连接
	db.DB().SetMaxOpenConns(100)                  // 设置连接池的最大连接
	db.DB().SetConnMaxLifetime(time.Second * 30)     // 设置连接的最大存活时间

	DB = db
}
