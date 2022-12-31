package db

import (
	"go_category/configs"
	"log"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MySqlDB *gorm.DB
var PgSqlDB *gorm.DB

var RedisClient *redis.Client

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

// InitDB 初始化db
func InitDB() (err error) {
	initMysql()
	initPgSql()
	initRedis()
	return dbInitError
}

// CloseDB 关闭数据库连接
func CloseDB() {
	closeMysql()
	closePgsql()
	closeRedis()
}

// Automigrate 自动修改库表结构
func Automigrate() {
	mysqlAutoMigrate()
	pgsqlAutoMigrate()
}
