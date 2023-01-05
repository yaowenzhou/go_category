package configs

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// DbType TODO
type DbType string

const (
	// MYSQL_TYPE TODO
	MYSQL_TYPE DbType = "mysql"
	// PGSQL_TYPE TODO
	PGSQL_TYPE DbType = "postgresql"
	// REDIS_TYPE TODO
	REDIS_TYPE DbType = "redis"
)

// ConfigSrcType TODO
type ConfigSrcType int

const (
	// CONSUL_JSON TODO
	CONSUL_JSON ConfigSrcType = iota
	// LOCAL_YAML TODO
	LOCAL_YAML
)

// LogFile TODO
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

// GetDbConfigFromLocalYaml 从本地yaml配置文件读取db配置
func GetDbConfigFromLocalYaml(dbType DbType) (dbConfig *DbConfig, err error) {
	dbConfig = &DbConfig{}
	// TODO: viper明显有bug，在制定了config_type的情况下，还是以默认支持的扩展名列表进行搜索
	// TODO: 因此比如这里，我已经设置了viper.SetConfigType("yaml")
	// TODO: 但是它找到config.json就不找了，就很坑
	// TODO: 所以此处实际上使用的是config.json
	viper := viper.New()
	viper.SetConfigName("config") // 配置文件名字，注意没有扩展名
	// 如果配置文件的名称中没有包含扩展名，那么该字段是必需的
	// 此函数设置的文件类型不在'viper.SupportedExts'中定义的时候
	viper.SetConfigType("yaml")
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
	case LOCAL_YAML:
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
