package db

import (
	"context"
	"fmt"
	"go_category/configs"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MySqlDB *gorm.DB
var PgSqlDB *gorm.DB

var RedisClient *redis.Client

var dbInitError error

// genGormConfig 生成gorm.Config
func genGormConfig(dbType configs.DbType) (cfg *gorm.Config) {
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
	ns := schema.NamingStrategy{SingularTable: true}
	if dbType == configs.PGSQL_TYPE {
		ns.TablePrefix = "tenant_wenzhouyao." // schema以Tname为前缀
	}
	return &gorm.Config{
		Logger:         newLogger,
		NamingStrategy: ns,
	}
}

// InitDB 初始化db
func InitDB(ctx context.Context, dbs []configs.DbType) (err error) {
	typeMap := make(map[configs.DbType]struct{})
	for _, v := range dbs {
		if _, ok := typeMap[v]; ok { // 不允许重复初始化
			CloseDB()
			errMsg := fmt.Sprintf("%s has been inited", v)
			return errors.Wrap(errors.New(errMsg), errMsg+"!")
		}
		typeMap[v] = struct{}{}
		switch v {
		case configs.MYSQL_TYPE:
			initMysql()
		case configs.PGSQL_TYPE:
			initPgSql()
		case configs.REDIS_TYPE:
			initRedis(ctx)
		}
	}
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
