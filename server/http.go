package server

import (
	"fmt"
	"strconv"

	"github.com/ztalab/ZAManager/pkg/confer"
	"github.com/ztalab/ZAManager/pkg/gin"
	"github.com/ztalab/ZAManager/pkg/middle"
	"github.com/ztalab/ZAManager/route"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunHTTP() {
	engine := gin.NewGin()
	//engine.Use(middle.RequestID())
	// 仅针对开发环境生效的组件
	//if confer.ConfigEnvIsDev() {
	// 跨域中间件
	engine.Use(middle.CorsV2())
	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}
	engine.Use(middle.Session("ZTA_session"))
	route.Home(engine)
	route.Api(engine)
	route.NotFound(engine)
	httpPort := confer.ConfigAppGetInt("port", 80)
	portStr := ":" + strconv.Itoa(httpPort)
	fmt.Println("start", httpPort)
	gin.ListenHTTP(portStr, engine, 10)
}
