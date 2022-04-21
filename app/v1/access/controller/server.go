package v1

import (
	"strconv"

	"github.com/ztalab/ZAManager/app/base/controller"
	"github.com/ztalab/ZAManager/app/v1/access/model/mparam"
	"github.com/ztalab/ZAManager/app/v1/access/service"
	"github.com/ztalab/ZAManager/pconst"
	"github.com/ztalab/ZAManager/pkg/response"

	"github.com/gin-gonic/gin"
)

// @Summary GetServer
// @Description 获取ZTA的server
// @Tags ZTA
// @Produce  json
// @Success 200 {object} controller.Res
// @Router /access/server [get]
func GetServer(c *gin.Context) {
	param := mparam.GetServer{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.GetServer(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary AddServer
// @Description 新增ZTA的server
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Server body mparam.AddServer true "新增ZTA的server"
// @Success 200 {object} controller.Res
// @Router /access/server [post]
func AddServer(c *gin.Context) {
	param := &mparam.AddServer{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.AddServer(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary EditServer
// @Description 修改ZTA的server
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Server body mparam.EditServer true "修改ZTA的server"
// @Success 200 {object} controller.Res
// @Router /access/server [put]
func EditServer(c *gin.Context) {
	param := &mparam.EditServer{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code = service.EditServer(c, param)
	response.UtilResponseReturnJson(c, code, nil)
}

// @Summary DelServer
// @Description 删除ZTA的server
// @Tags ZTA
// @Produce  json
// @Param id path int true "主键ID"
// @Success 200 {object} controller.Res
// @Router /access/server/{id} [delete]
func DelServer(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.UtilResponseReturnJsonFailed(c, pconst.CODE_COMMON_PARAMS_INCOMPLETE)
		return
	}
	code := service.DelServer(c, uint64(idInt))
	response.UtilResponseReturnJson(c, code, nil)
}
