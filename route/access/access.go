package access

import (
	v1 "github.com/ztalab/ZAManager/app/v1/access/controller"

	"github.com/gin-gonic/gin"
)

func APIAccess(parentRoute gin.IRouter) {
	//r := parentRoute.Group("access", middle.Auth())
	r := parentRoute.Group("access")
	{
		resource := r.Group("resource")
		{
			resource.GET("", v1.GetResource)
			resource.POST("", v1.AddResource)
			resource.PUT("", v1.EditResource)
			resource.DELETE("/:id", v1.DelResource)
		}
		relay := r.Group("relay")
		{
			relay.GET("", v1.GetRelay)
			relay.POST("", v1.AddRelay)
			relay.PUT("", v1.EditRelay)
			relay.DELETE("/:id", v1.DelRelay)
		}
		server := r.Group("server")
		{
			server.GET("", v1.GetServer)
			server.POST("", v1.AddServer)
			server.PUT("", v1.EditServer)
			server.DELETE("/:id", v1.DelServer)
		}
		client := r.Group("client")
		{
			client.GET("", v1.GetClient)
			client.POST("", v1.AddClient)
			client.PUT("", v1.EditClient)
			client.DELETE("/:id", v1.DelClient)
		}
	}

}
