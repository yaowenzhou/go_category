package main

import (
	"context"
	"fmt"
	"go_category/configs"
	"go_category/domain/db"
	"log"
)

func setLogFile() {
	log.SetOutput(configs.LogFile)
}

func main() {
	ctx := context.Background()
	setLogFile()
	defer configs.CloseLogFile()
	err := db.InitDB(ctx, []configs.DbType{configs.MYSQL_TYPE, configs.PGSQL_TYPE, configs.REDIS_TYPE})
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()
	db.Automigrate()
	r, err := db.RedisClient.Incr(ctx, "wenzhouyao_use_cnt").Result()
	if err != nil {
		fmt.Println("db.RedisClient.Incr fail, err:", err.Error())
		return
	}
	fmt.Println("db.RedisClient.Incr reply:", r)
	cnt, err := db.RedisClient.Get(ctx, "wenzhouyao_use_cnt").Int()
	if err != nil {
		fmt.Println("db.RedisClient.Get fail, err:", err.Error())
	} else {
		fmt.Println("wenzhouyao_use_cnt:", cnt)
	}
}
