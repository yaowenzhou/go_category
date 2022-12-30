package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type DbConfig struct {
	UserName string
	Host     string
	Password string
	Port     int
	DbName   string
	TimeOut  int
}

type DbType string

const MYSQL DbType = "MySQL"
const PGSQL DbType = "PostgreSQL"

func GetConfig(dbType DbType) (dbConfig *DbConfig, err error) {
	dbConfig := &DbConfig{}
	viper.SetConfigName("config")     // 配置文件名字，注意没有扩展名
	viper.SetConfigType("yaml")       // 如果配置文件的名称中没有包含扩展名，那么该字段是必需的
	viper.AddConfigPath("./configs/") // 配置文件的路径
	err = viper.ReadInConfig()        // 查找并读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if dbType != MYSQL && dbType != PGSQL {
		return nil, errors.Wrap("dbType must be 'MySQL' or 'PostgreSQL'")
	}
	dbConfig.Host = viper.GetString(fmt.Sprintf("%s.ip", dbType))
	return &DbConfig
}
