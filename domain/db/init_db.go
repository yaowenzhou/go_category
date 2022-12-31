package db

import (
	"fmt"
	"go_category/configs"
	"go_category/domain/model"
	"log"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MySqlDB *gorm.DB
var PgSqlDB *gorm.DB

// var RedisClient

var dbInitError error

// genGormConfig 生成gorm.Config
func genGormConfig() (cfg *gorm.Config) {
	if dbInitError != nil {
		return
	}
	newLogger := logger.New(
		log.New(configs.LogFile, "\r\n", log.LstdFlags), // 指定日志文件
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值, 200ms
			LogLevel:                  logger.Info,            // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,                   // 彩色打印
		},
	)
	return &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁止复数表名
		},
	}
}

// initMysql 初始化mysql
func initMysql() {
	if dbInitError != nil {
		return
	}
	var dbConfig *configs.DbConfig
	dbConfig, dbInitError = configs.GetDbConfig(configs.MYSQL)
	if dbInitError != nil {
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds",
		dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.TimeOut)
	log.Printf("mysql dsn: %s\n", dsn)
	MySqlDB, dbInitError = gorm.Open(mysql.Open(dsn), genGormConfig())
	if dbInitError != nil {
		dbInitError = errors.Wrap(dbInitError, "gorm.Open fail.")
		return
	}
	sqlDb, err := MySqlDB.DB()
	if err != nil {
		dbInitError = errors.Wrap(dbInitError, "get sqldb fail.")
		return
	}
	sqlDb.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDb.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

// initPgSql 初始化pgsql
func initPgSql() {
	if dbInitError != nil {
		return
	}
	var dbConfig *configs.DbConfig
	dbConfig, dbInitError = configs.GetDbConfig(configs.PGSQL)
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
	sqlDb.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDb.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

func initRedis() {

}

// InitDB 初始化db
func InitDB() (err error) {
	initMysql()
	initPgSql()
	return dbInitError
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if MySqlDB != nil {
		sqlDb, err := MySqlDB.DB()
		if err == nil {
			sqlDb.Close()
		}
	}
	if PgSqlDB != nil {
		sqlDb, err := PgSqlDB.DB()
		if err == nil {
			sqlDb.Close()
		}
	}
}

func Automigrate() {
	MySqlDB.AutoMigrate(&model.User{})
}
