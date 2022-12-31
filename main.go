package main

import (
	"context"
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
}
