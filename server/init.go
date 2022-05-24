package server

import (
	"github.com/urfave/cli"
	"github.com/ztalab/ZAManager/pkg/confer"
	"github.com/ztalab/ZAManager/pkg/logger"
	"github.com/ztalab/ZAManager/pkg/longpoll"
	"github.com/ztalab/ZAManager/pkg/mysql"
	"github.com/ztalab/ZAManager/pkg/redis"
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
	if err = longpoll.Init(); err != nil {
		return
	}
	if err = redis.Init(&cfg.Redis); err != nil {
		return
	}
	if err = mysql.Init(&cfg.Mysql); err != nil {
		return
	}
	if err = mysql.SqlMigrate(); err != nil {
		return
	}
	return
}
