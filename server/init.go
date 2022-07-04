package server

import (
	"github.com/urfave/cli"
	"github.com/ztalab/cloudslit/pkg/confer"
	"github.com/ztalab/cloudslit/pkg/logger"
	"github.com/ztalab/cloudslit/pkg/mysql"
	"github.com/ztalab/cloudslit/pkg/redis"
	"github.com/ztalab/cloudslit/pkg/web3/eth"
	"github.com/ztalab/cloudslit/pkg/web3/w3s"
)

func InitService(c *cli.Context) (err error) {
	if err = confer.Init(c.String("c")); err != nil {
		return
	}
	cfg := confer.GlobalConfig()
	logger.Init(&logger.Config{
		Level:       logger.LogLevel(),
		Filename:    logger.LogFile(),
		SendToFile:  logger.SendLogToFile(),
		Development: confer.ConfigEnvIsDev(),
	})
	if err = redis.Init(&cfg.Redis); err != nil {
		return
	}
	if err = mysql.Init(&cfg.Mysql); err != nil {
		return
	}
	if err = w3s.Init(&cfg.Web3); err != nil {
		return
	}
	if err = eth.Init(&cfg.Web3); err != nil {
		return
	}
	//if err = mysql.SqlMigrate(); err != nil {
	//	return
	//}
	// 判断是否开启P2P
	if confer.GlobalConfig().P2P.Enable {
		return runP2P()
	}
	return
}
