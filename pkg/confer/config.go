package confer

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var globalConfig *ServerConfig
var mutex sync.RWMutex

type ServerConfig struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Log   Log   `mapstructure:"log" json:"log" yaml:"log"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	CA    CA    `mapstructure:"ca" json:"ca" yaml:"ca"`
	sync.RWMutex
}

type App map[string]interface{}

type Redis struct {
	Addr   string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}

type Mysql struct {
	DBName string   `mapstructure:"dbname" json:"dbName" yaml:"dbname"`
	Prefix string   `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Pool   DBPool   `mapstructure:"pool" json:"pool" yaml:"pool"`
	Write  DBBase   `mapstructure:"write" json:"write" yaml:"write"`
	Reads  []DBBase `mapstructure:"reads" json:"reads" yaml:"reads"`
}

type DBPool struct {
	PoolMinCap      int   `mapstructure:"pool-min-cap" json:"poolMinCap" yaml:"pool-min-cap"`
	PoolExCap       int   `mapstructure:"pool-ex-cap" json:"poolExCap" yaml:"pool-ex-cap"`
	PoolMaxCap      int   `mapstructure:"pool-max-cap" json:"pool-max-cap" yaml:"pool-max-cap"`
	PoolIdleTimeout int   `mapstructure:"pool-idle-timeout" json:"poolIdleTimeout" yaml:"pool-idle-timeout"`
	PoolWaitCount   int64 `mapstructure:"pool-wait-count" json:"poolWaitCount" yaml:"pool-wait-count"`
	PoolWaitTimeout int   `mapstructure:"pool-wai-timeout" json:"poolWaitTimeout" yaml:"pool-wai-timeout"`
}

type DBBase struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DBName   string `json:"-"`
	Prefix   string `json:"-"`
}

type CA struct {
	BaseURL string `mapstructure:"base-url" json:"base_url" yaml:"base-url"`
	SignURL string `mapstructure:"sign-url" json:"sign_url" yaml:"sign-url"`
	OcspURL string `mapstructure:"ocsp-url" json:"ocsp_url" yaml:"ocsp-url"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
	AuthKey string `mapstructure:"auth-key" json:"auth_key" yaml:"auth-key"`
}

type Log struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"`
	SendToFile  bool   `mapstructure:"send-to-file" json:"send_to_file" yaml:"send-to-file"`
	Filename    string `mapstructure:"filename" json:"filename" yaml:"filename"`
	NoCaller    bool   `mapstructure:"no-calle" json:"no_caller" yaml:"no-caller"`
	NoLogLevel  bool   `mapstructure:"no-log-level" json:"no_log_level" yaml:"no-log-level"`
	Development bool   `mapstructure:"development" json:"development" yaml:"development"`
	MaxSize     int    `mapstructure:"max-size" json:"max_size" yaml:"max-size"` // megabytes
	MaxAge      int    `mapstructure:"max-age" json:"max_age" yaml:"max-age"`    // days
	MaxBackups  int    `mapstructure:"max-backups" json:"max_backups" yaml:"max-backups"`
}

func Init(configURL string) (err error) {
	f, err := os.Open(configURL)
	if err != nil {
		return
	}
	if err = yaml.NewDecoder(f).Decode(&globalConfig); err != nil {
		return
	}
	handleConfig(globalConfig)
	return
}

func handleConfig(config *ServerConfig) {
	config.replaceByEnv(&config.Redis.Addr)
	config.replaceByEnv(&config.Mysql.Write.Host)
	config.replaceByEnv(&config.Mysql.Write.User)
	config.replaceByEnv(&config.Mysql.Write.Password)
	config.replaceByEnv(&config.CA.BaseURL)
	config.replaceByEnv(&config.CA.SignURL)
	config.replaceByEnv(&config.CA.OcspURL)
	config.replaceByEnv(&config.CA.AuthKey)
	config.Mysql.Write.DBName = globalConfig.Mysql.DBName
	config.Mysql.Write.Prefix = globalConfig.Mysql.Prefix
	return
}

func (*ServerConfig) replaceByEnv(conf *string) {
	if s := os.Getenv(*conf); len(s) > 0 {
		*conf = s
	}
}

func GlobalConfig() *ServerConfig {
	mutex.RLock()
	defer mutex.RUnlock()
	return globalConfig
}
