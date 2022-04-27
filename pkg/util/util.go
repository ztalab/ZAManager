package util

import (
	"encoding/json"
	"net"
	"regexp"

	"github.com/ztalab/ZAManager/app/v1/user/model/mmysql"

	"github.com/gin-gonic/gin"
)

func IsIP(ipv4 string) bool {
	if ip := net.ParseIP(ipv4); ip == nil {
		return false
	}
	return true
}

func IsCIDR(cidr string) bool {
	p := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`)
	return p.MatchString(cidr)
}

func GetCookieFromGin(ctx *gin.Context, key string) (value string) {
	value, _ = ctx.Cookie(key)
	return
}

func User(c *gin.Context) (user *mmysql.User) {
	if userBytes, ok := c.Get("user"); ok {
		if err := json.Unmarshal(userBytes.([]byte), &user); err != nil {
			return
		}
		return nil
	}
	return nil
}
