package configs

import (
	"fmt"

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

const MYSQL DbType = "mysql"
const PGSQL DbType = "postgresql"

// GetDbConfig 读取DB配置
func GetDbConfig(dbType DbType) (dbConfig *DbConfig, err error) {
	if dbType != MYSQL && dbType != PGSQL {
		errMsg := "dbType must be 'mysql' or 'postgresql'"
		return nil, errors.Wrap(errors.New(errMsg), errMsg+"!")
	}
	dbConfig = &DbConfig{}
	viper.SetConfigName("config")     // 配置文件名字，注意没有扩展名
	viper.SetConfigType("yaml")       // 如果配置文件的名称中没有包含扩展名，那么该字段是必需的
	viper.AddConfigPath("./configs/") // 配置文件的路径
	err = viper.ReadInConfig()        // 查找并读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	dbConfig.UserName = viper.GetString(fmt.Sprintf("%s.user", dbType))
	dbConfig.Host = viper.GetString(fmt.Sprintf("%s.ip", dbType))
	dbConfig.Password = viper.GetString(fmt.Sprintf("%s.password", dbType))
	dbConfig.Port = viper.GetInt(fmt.Sprintf("%s.port", dbType))
	dbConfig.DbName = viper.GetString(fmt.Sprintf("%s.database", dbType))
	dbConfig.TimeOut = viper.GetInt(fmt.Sprintf("%s.timeout", dbType))
	return dbConfig, nil
}
