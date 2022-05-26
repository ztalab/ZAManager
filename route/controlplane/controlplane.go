package controlplane

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/ztalab/ZAManager/app/v1/controlplane/controller"
)

func APIControlPlane(parentRoute gin.IRouter) {
	controlplane := parentRoute.Group("controlplane")
	{
		controlplane.GET("/machine/:machine_id", v1.LoginUrl)
		//controlplane.POST("/machine/auth/pub", wrapWithContext(longpoll.Manger().PublishHandler))
		//controlplane.GET("/machine/auth/poll", wrapWithContext(longpoll.Manger().SubscriptionHandler))
		controlplane.GET("/machine/auth/poll", v1.MachineLongPoll)
	}
}

func wrapWithContext(lpHandler func(http.ResponseWriter, *http.Request)) func(*gin.Context) {
	return func(c *gin.Context) {
		lpHandler(c.Writer, c.Request)
	}
}
