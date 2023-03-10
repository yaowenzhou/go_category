package test

import (
	"context"
	"fmt"
	"go_category/configs"
	"go_category/domain/db"
	"log"
	"testing"
)

func test(ctx context.Context, t *testing.T, dbType configs.DbType, f func()) {
	configs.InitConsulConfig("192.168.1.99", 8500, "/category/config/")
	log.SetOutput(configs.LogFile)
	defer configs.CloseLogFile()
	err := db.InitDB(ctx, []configs.DbType{dbType}, configs.CONSUL_JSON)
	if err != nil {
		fmt.Println("db.InitDB fail:", err)
		return
	}
	defer db.CloseDB()
	f()
}
