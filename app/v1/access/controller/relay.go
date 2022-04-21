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

// @Summary GetRelay
// @Description 获取ZTA的relay
// @Tags ZTA
// @Produce  json
// @Success 200 {object} controller.Res
// @Router /access/relay [get]
func GetRelay(c *gin.Context) {
	param := mparam.GetRelay{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.GetRelay(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary AddRelay
// @Description 新增ZTA的relay
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Relay body mparam.AddRelay true "新增ZTA的relay"
// @Success 200 {object} controller.Res
// @Router /access/relay [post]
func AddRelay(c *gin.Context) {
	param := &mparam.AddRelay{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.AddRelay(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary EditRelay
// @Description 修改ZTA的relay
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Relay body mparam.EditRelay true "修改ZTA的relay"
// @Success 200 {object} controller.Res
// @Router /access/relay [put]
func EditRelay(c *gin.Context) {
	param := &mparam.EditRelay{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code = service.EditRelay(c, param)
	response.UtilResponseReturnJson(c, code, nil)
}

// @Summary DelRelay
// @Description 删除ZTA的relay
// @Tags ZTA
// @Produce  json
// @Param id path int true "主键ID"
// @Success 200 {object} controller.Res
// @Router /access/relay/{id} [delete]
func DelRelay(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.UtilResponseReturnJsonFailed(c, pconst.CODE_COMMON_PARAMS_INCOMPLETE)
		return
	}
	code := service.DelRelay(c, uint64(idInt))
	response.UtilResponseReturnJson(c, code, nil)
}
