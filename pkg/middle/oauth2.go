package middle

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Oauth2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if user := session.Get("user"); user != nil {
			ctx.Set("user", user)
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
