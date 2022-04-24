package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ztalab/ZAManager/app/v1/user/controller"
	"github.com/ztalab/ZAManager/pkg/middle"
)

func APIUser(parentRoute gin.IRouter) {
	user := parentRoute.Group("user")
	{
		user.GET("/login/:company", v1.Login)
		user.GET("/oauth2/callback/:company", v1.Oauth2Callback)
		user.GET("/detail", middle.Oauth2(), v1.UserDetail)
	}
}
