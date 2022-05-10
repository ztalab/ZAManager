package v1

import (
	"github.com/ztalab/ZAManager/app/base/controller"
	"github.com/ztalab/ZAManager/app/v1/access/model/mparam"
	"github.com/ztalab/ZAManager/app/v1/access/service"
	"github.com/ztalab/ZAManager/pkg/response"

	"github.com/gin-gonic/gin"
)

// @Summary ClientList
// @Description 获取ZTA的client
// @Tags ZTA
// @Produce  json
// @Success 200 {object} controller.Res
// @Router /access/client [get]
func ClientList(c *gin.Context) {
	param := mparam.ClientList{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.ClientList(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary AddClient
// @Description 新增ZTA的client
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Client body mparam.AddClient true "新增ZTA的client"
// @Success 200 {object} controller.Res
// @Router /access/client [post]
func AddClient(c *gin.Context) {
	param := &mparam.AddClient{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.AddClient(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary EditClient
// @Description 修改ZTA的client
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Client body mparam.EditClient true "修改ZTA的client"
// @Success 200 {object} controller.Res
// @Router /access/client [put]
func EditClient(c *gin.Context) {
	param := &mparam.EditClient{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code = service.EditClient(c, param)
	response.UtilResponseReturnJson(c, code, nil)
}

// @Summary DelClient
// @Description 删除ZTA的client
// @Tags ZTA
// @Produce  json
// @Param uuid path string true "uuid"
// @Success 200 {object} controller.Res
// @Router /access/client/{uuid} [delete]
func DelClient(c *gin.Context) {
	code := service.DelClient(c, c.Param("uuid"))
	response.UtilResponseReturnJson(c, code, nil)
}
