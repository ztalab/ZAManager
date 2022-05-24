package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/ztalab/ZAManager/app/v1/controlplane/dao/redis"

	"github.com/ztalab/ZAManager/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/ztalab/ZAManager/pconst"
	"github.com/ztalab/ZAManager/pkg/confer"
)

func GetLoginUrl(c *gin.Context, machine string) (code int, loginURL string) {
	// 通过machineID和当前时间戳，计算出唯一的hash，作为登陆的地址path
	hash := util.NewMd5(fmt.Sprintf("%s%d", machine, time.Now().UnixNano()))
	// hash 放入redis缓存
	err := redis.NewMachine(c).SetLoginHash(machine, hash)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	loginURL = fmt.Sprintf("%s/a/%s", confer.ConfigAppGet("domain"), hash)
	return
}

func MachineOauth(c *gin.Context, hash string) {
	// 判断当前hash是否存在或者是否在有消息内
	exist, _, err := redis.NewMachine(c).GetLoginHash(hash)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("oauth error"))
		return
	}
	if exist {
		session := sessions.Default(c)
		session.Set("machine", hash)
		session.Save()
		c.Redirect(http.StatusSeeOther, "/api/v1/user/login/github")
	} else {
		// TODO 重定向到404页面
		c.String(http.StatusNotFound, "auth key not exist or expired")
		return
	}

}
