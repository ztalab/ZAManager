package middle

import (
	"encoding/json"
	"net/http"

	"github.com/ztalab/ZAManager/app/v1/user/model/mmysql"

	"github.com/ztalab/ZAManager/pkg/confer"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Oauth2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if user := session.Get("user"); user != nil {
			ctx.Set("user", user)
			ctx.Next()
		} else {
			if confer.ConfigEnvGet() == "dev" {
				userBytes, _ := json.Marshal(&mmysql.User{
					Email: "nisainan@github.com",
					UUID:  "3933d404-2025-4851-bfe3-1c07c5280c72",
				})
				ctx.Set("user", userBytes)
				ctx.Next()
			} else {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}
