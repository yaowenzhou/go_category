package configs

import (
	"fmt"

	cfgConsul "github.com/go-micro/plugins/v4/config/source/consul"
	regConsul "github.com/go-micro/plugins/v4/registry/consul"
	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/registry"
)

var consulConfig config.Config

// GetConsulRegistry 获取consul配置
func GetConsulRegistry(host string, port int64) registry.Registry {
	consulReg := regConsul.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%d", host, port)),
	)
	return consulReg
}

// InitConsulConfig 设置配置中心
func InitConsulConfig(host string, port int64, prefix string) (err error) {
	consulSource := cfgConsul.NewSource(
		// optionally specify consul address; default to localhost:8500
		cfgConsul.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		// optionally specify prefix; defaults to /micro/config
		cfgConsul.WithPrefix(prefix),
		// optionally strip the provided prefix from the keys, defaults to false
		cfgConsul.StripPrefix(true),
	)
	// 配置初始化
	consulConfig, err = config.NewConfig()
	if err != nil {
		return errors.Wrap(err, "config.NewConfig fail")
	}
	// 加载配置
	err = consulConfig.Load(consulSource)
	if err != nil {
		err = errors.Wrap(err, "config.Load fail")
	}
	return
}

// GetDbConfigFromConsul 从consul获取db的配置
func GetDbConfigFromConsul(path ...string) (dbConfig *DbConfig, err error) {
	dbConfig = &DbConfig{}
	err = consulConfig.Get(path...).Scan(dbConfig)
	if err != nil {
		return nil, errors.Wrap(err, "scan from consul fail.")
	}
	return
}
