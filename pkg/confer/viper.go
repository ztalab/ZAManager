package confer

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var globalConfig *ServerConfig
var mutex sync.RWMutex

func Init(configURL string) (err error) {
	v := viper.New()
	v.SetConfigFile(configURL)
	err = v.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("fatal error config file: %w", err)
		return
	}
	if err = v.Unmarshal(&globalConfig); err != nil {
		return
	}
	handleConfig(globalConfig)
	return
}

func handleConfig(config *ServerConfig) {
	config.Mysql.Write.DBName = globalConfig.Mysql.DBName
	config.Mysql.Write.Prefix = globalConfig.Mysql.Prefix

	return
}

func GlobalConfig() *ServerConfig {
	mutex.RLock()
	defer mutex.RUnlock()
	return globalConfig
}
