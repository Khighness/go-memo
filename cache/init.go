package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

// @Author KHighness
// @Update 2022-02-13

// RedisClient
var (
	RedisClient  *redis.Client
	RedisAddr    string
	RedisAuth    string
	RedisDb      int
)

// 初始化Redis
func Redis(file *ini.File) {
	loadRedisConfig(file)
	connectRedis()
}

// 读取Redis配置
func loadRedisConfig(file *ini.File)  {
	redis := file.Section("redis")
	RedisAddr = redis.Key("RedisAddr").String()
	RedisAuth = redis.Key("RedisAuth").String()
	Db, err := redis.Key("RedisDb").Int()
	if err != nil {
		panic(err)
	} else {
		RedisDb = Db
	}
}

// 连接到Redis
func connectRedis()  {
	client := redis.NewClient(&redis.Options{
		Addr:         RedisAddr,
		Password:     RedisAuth,
		DB:           RedisDb,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
		PoolSize:     20,
		MinIdleConns: 1,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to redis server [%s] error", RedisAuth))
		panic(err)
	}

	RedisClient = client
}
