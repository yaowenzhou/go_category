package main

import (
	"context"
	"go_category/configs"
	"go_category/domain/db"
	"go_category/handler"
	categoryPb "go_category/proto/pb/category"
	"log"

	"github.com/sirupsen/logrus"
	"go-micro.dev/v4"
)

func setLogFile() {
	log.SetOutput(configs.LogFile)
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(configs.LogFile)
	ctx := context.Background()
	setLogFile()
	defer configs.CloseLogFile()
	configs.InitConsulConfig("192.168.1.99", 8500, "/category/config/")
	err := db.InitDB(ctx,
		[]configs.DbType{
			configs.MYSQL_TYPE,
			configs.PGSQL_TYPE,
			configs.REDIS_TYPE,
		},
		configs.LOCAL_YAML,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()
	db.Automigrate()
	// server := server.NewServer()
	service := micro.NewService(
		micro.Name("category"),
		micro.Version("lateest"),
		micro.Address("192.168.1.99:8080"),
		// 添加consul作为注册中心
		micro.Registry(configs.GetConsulRegistry("192.168.1.99", 8500)),
	)
	service.Init()
	categoryPb.RegisterCategoryHandler(service.Server(), new(handler.Category))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
