package main

import (
	"fmt"
	"go-memo/cache"
	"go-memo/conf"
)

// @Author KHighness
// @Update 2022-02-13
// @Swagger http://localhost:3333/swagger/index.html
func main() {
	conf.Init()
	//cache.RedisClient.HSet("go-memo", "author", "Khighness")
	//cache.RedisClient.HSet("go-memo", "version", "1.0.0")
	fmt.Println(cache.RedisClient.HGet("go-memo", "author"))
}