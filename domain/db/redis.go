package db

import (
	"fmt"
	"go_category/configs"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// initRedis 初始化redis客户端
func initRedis() {
	if dbInitError != nil {
		return
	}
	var dbConfig *configs.DbConfig
	dbConfig, dbInitError = configs.GetDbConfig(configs.REDIS_TYPE)
	if dbInitError != nil {
		return
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port), // 连接地址
		Password: dbConfig.Password,                                  // 连接密码
		DB:       cast.ToInt(dbConfig.DbName),                        // 默认连接库
		PoolSize: dbConfig.MaxOpenConns,                              // 连接池大小
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		dbInitError = errors.Wrap(err, "redis connect fail.")
	}
}

// closeRedis 关闭redis句柄
func closeRedis() {
	if RedisClient != nil {
		RedisClient.Close()
	}
}
