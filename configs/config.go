package configs

import (
	"fmt"
	"go_category/consts"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// DbType TODO
type DbType string

// MYSQL_TYPE TODO
const MYSQL_TYPE DbType = "mysql"

// PGSQL_TYPE TODO
const PGSQL_TYPE DbType = "postgresql"

// REDIS_TYPE
const REDIS_TYPE DbType = "redis"

type ConfigSrcType int

const (
	CONSUL_JSON ConfigSrcType = iota
	LOCAL_FILE
)

// LogFile 日志文件
var LogFile *os.File

// DbConfig 配置
type DbConfig struct {
	UserName     string `json:"user"`
	Host         string `json:"ip"`
	Password     string `json:"password"`
	Port         int    `json:"port"`
	DbName       string `json:"database"`
	TimeOut      int    `json:"timeout"`        // 超时时间
	MaxOpenConns int    `json:"max_open_conns"` // 最大连接数量
	MaxIdleConns int    `json:"max_idle_conns"` // 最大空闲连接数量
}

func init() {
	var err error
	LogFile, err = os.OpenFile(consts.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		panic(fmt.Errorf("open config file: %s", err.Error()))
	}
}

// GetDbConfigFromLocalYaml 从本地yaml配置文件读取db配置
func GetDbConfigFromLocalYaml(dbType DbType) (dbConfig *DbConfig, err error) {
	dbConfig = &DbConfig{}
	viper.SetConfigName("config") // 配置文件名字，注意没有扩展名
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有包含扩展名，那么该字段是必需的
	// /home/adc/projects/go_category/configs/config.go
	// viper.AddConfigPath("./configs/") // 配置文件的路径
	viper.AddConfigPath("/home/adc/projects/go_category/configs/")
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig fail")
	}
	dbConfig.UserName = viper.GetString(fmt.Sprintf("%s.user", dbType))
	dbConfig.Host = viper.GetString(fmt.Sprintf("%s.ip", dbType))
	dbConfig.Password = viper.GetString(fmt.Sprintf("%s.password", dbType))
	dbConfig.Port = viper.GetInt(fmt.Sprintf("%s.port", dbType))
	dbConfig.DbName = viper.GetString(fmt.Sprintf("%s.database", dbType))
	dbConfig.TimeOut = viper.GetInt(fmt.Sprintf("%s.timeout", dbType))
	dbConfig.MaxOpenConns = viper.GetInt(fmt.Sprintf("%s.max_open_conns", dbType))
	dbConfig.MaxIdleConns = viper.GetInt(fmt.Sprintf("%s.max_idle_conns", dbType))
	return dbConfig, nil
}

// GetDbConfig 读取DB配置
func GetDbConfig(dbType DbType, cst ConfigSrcType) (dbConfig *DbConfig, err error) {
	if dbType != MYSQL_TYPE && dbType != PGSQL_TYPE && dbType != REDIS_TYPE {
		errMsg := "dbType must be one of 'mysql|postgresql|redis'"
		return nil, errors.Wrap(errors.New(errMsg), errMsg+"!")
	}
	switch cst {
	case CONSUL_JSON:
		return GetDbConfigFromConsul(string(dbType))
	case LOCAL_FILE:
		return GetDbConfigFromLocalYaml(dbType)
	default:
		errMsg := fmt.Sprintf("wrong param of dbType(%s)", dbType)
		return nil, errors.Wrap(errors.New(errMsg), errMsg)
	}
}

// GetLogFile 获取日志文件句柄
func GetLogFile() *os.File {
	return LogFile
}

// CloseLogFile 关闭日志文件
func CloseLogFile() {
	LogFile.Close()
}
