package db

import (
	"fmt"
	"go_category/configs"
	"go_category/domain/model"
	"log"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// initMysql 初始化mysql
func initMysql(cst configs.ConfigSrcType) {
	if dbInitError != nil {
		return
	}
	var dbConfig *configs.DbConfig
	dbConfig, dbInitError = configs.GetDbConfig(configs.MYSQL_TYPE, cst)
	if dbInitError != nil {
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds",
		dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.TimeOut)
	log.Printf("mysql dsn: %s\n", dsn)
	MySqlDB, dbInitError = gorm.Open(mysql.Open(dsn), genGormConfig(configs.MYSQL_TYPE))
	if dbInitError != nil {
		dbInitError = errors.Wrap(dbInitError, "gorm.Open fail.")
		return
	}
	sqlDb, err := MySqlDB.DB()
	if err != nil {
		dbInitError = errors.Wrap(dbInitError, "get sqldb fail.")
		return
	}
	// 设置数据库连接池最大连接数
	sqlDb.SetMaxOpenConns(dbConfig.MaxOpenConns)
	// 连接池最大允许的空闲连接数
	// 如果没有sql任务需要执行的连接数大于此值，超过的连接会被连接池关闭。
	sqlDb.SetMaxIdleConns(dbConfig.MaxIdleConns)
}

// closeMysql 关闭mysql句柄
func closeMysql() {
	if MySqlDB != nil {
		sqlDb, err := MySqlDB.DB()
		if err == nil {
			sqlDb.Close()
		}
	}
}

// mysqlAutoMigrate 自动修改库表结构(mysql)
func mysqlAutoMigrate() {
	MySqlDB.AutoMigrate(
		&model.User{},
		&model.Category{},
	)
}
