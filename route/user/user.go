package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	v1 "github.com/ztalab/ZAManager/app/v1/user/controller"
	"github.com/ztalab/ZAManager/pkg/middle"
)

func APIUser(parentRoute gin.IRouter) {
	user := parentRoute.Group("user", sessions.Sessions("session", cookie.NewStore([]byte("secret"))))
	{
		user.GET("/login/:company", v1.Login)
		user.GET("/oauth2/callback/:company", v1.Oauth2Callback)
		user.GET("/detail", middle.Oauth2(), v1.UserDetail)
	}
}
