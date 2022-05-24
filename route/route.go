package route

import (
	"github.com/ztalab/ZAManager/app/base/controller"
	v1 "github.com/ztalab/ZAManager/app/v1/controlplane/controller"
	"github.com/ztalab/ZAManager/pconst"
	"github.com/ztalab/ZAManager/pkg/confer"
	"github.com/ztalab/ZAManager/route/access"
	"github.com/ztalab/ZAManager/route/controlplane"
	"github.com/ztalab/ZAManager/route/system"
	"github.com/ztalab/ZAManager/route/user"

	"github.com/gin-gonic/gin"
)

// Home 主页
func Home(engine *gin.Engine) {
	engine.GET("", controller.Welcome)
}

func Api(engine *gin.Engine) {
	engine.GET("/a/:hash", v1.MachineOauth)
	prefix := confer.ConfigAppGetString("UrlPrefix", "")
	RouteV1 := engine.Group(prefix + pconst.APIAPIV1URL)
	{
		access.APIAccess(RouteV1)
		controlplane.APIControlPlane(RouteV1)
		system.APISystem(RouteV1)
		user.APIUser(RouteV1)
	}
}

func NotFound(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(404, "404 Not Found")
	})
}
