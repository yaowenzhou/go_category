package db

import (
	"fmt"
	"go_category/configs"
	"go_category/domain/model"
	"log"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initPgSql 初始化pgsql
func initPgSql() {
	if dbInitError != nil {
		return
	}
	var dbConfig *configs.DbConfig
	dbConfig, dbInitError = configs.GetDbConfig(configs.PGSQL_TYPE)
	if dbInitError != nil {
		return
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host, dbConfig.UserName, dbConfig.Password, dbConfig.DbName, dbConfig.Port)
	log.Printf("pgsql dsn: %s\n", dsn)
	PgSqlDB, dbInitError = gorm.Open(postgres.Open(dsn), genGormConfig())
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

// closePgsql 关闭pgsql句柄
func closePgsql() {
	if PgSqlDB != nil {
		sqlDb, err := PgSqlDB.DB()
		if err == nil {
			sqlDb.Close()
		}
	}
}

// pgsqlAutoMigrate 自动修改库表结构(pgsql)
func pgsqlAutoMigrate() {
	PgSqlDB.AutoMigrate(&model.User{})
}
