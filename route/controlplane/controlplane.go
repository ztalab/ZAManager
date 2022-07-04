package controlplane

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ztalab/cloudslit/app/v1/controlplane/controller"
)

func APIControlPlane(parentRoute gin.IRouter) {
	controlplane := parentRoute.Group("controlplane")
	{
		controlplane.GET("/machine/:machine_id", v1.LoginUrl)
		controlplane.GET("/machine/auth/poll", v1.MachineLongPoll)
	}
}
