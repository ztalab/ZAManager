package v1

import (
	"strconv"
	"strings"

	"github.com/ztalab/ZAManager/app/base/controller"
	"github.com/ztalab/ZAManager/app/v1/access/model/mparam"
	"github.com/ztalab/ZAManager/app/v1/access/service"
	"github.com/ztalab/ZAManager/pconst"
	"github.com/ztalab/ZAManager/pkg/response"
	"github.com/ztalab/ZAManager/pkg/util"

	"github.com/gin-gonic/gin"
)

// @Summary ResourceList
// @Description 获取ZTA的resource
// @Tags ZTA
// @Produce  json
// @Success 200 {object} controller.Res
// @Router /access/resource [get]
func ResourceList(c *gin.Context) {
	param := mparam.ResourceList{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	param.UserUUID = util.User(c).UUID
	code, data := service.ResourceList(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary AddResource
// @Description 新增ZTA的resource
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Resource body mparam.AddResource true "新增ZTA的resource"
// @Success 200 {object} controller.Res
// @Router /access/resource [post]
func AddResource(c *gin.Context) {
	param := &mparam.AddResource{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	if len(param.Type) > 0 && param.Type == "cidr" {
		// 判断是不是纯IP格式
		if strings.Contains(param.Host, "/") {
			if !util.IsCIDR(param.Host) {
				response.UtilResponseReturnJsonFailed(c, pconst.CODE_COMMON_PARAMS_INCOMPLETE)
				return
			}
		} else {
			if !util.IsIP(param.Host) {
				response.UtilResponseReturnJsonFailed(c, pconst.CODE_COMMON_PARAMS_INCOMPLETE)
				return
			}
		}
	}
	param.UserUUID = util.User(c).UUID
	code, data := service.AddResource(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary EditResource
// @Description 修改ZTA的resource
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Resource body mparam.EditResource true "修改ZTA的resource"
// @Success 200 {object} controller.Res
// @Router /access/resource [put]
func EditResource(c *gin.Context) {
	param := &mparam.EditResource{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	if len(param.Type) > 0 && param.Type == "cidr" {
		// 判断是不是纯IP格式
		if strings.Contains(param.Host, "/") {
			if !util.IsCIDR(param.Host) {
				response.UtilResponseReturnJsonFailed(c, pconst.CODE_COMMON_PARAMS_INCOMPLETE)
				return
			}
		} else {
			if !util.IsIP(param.Host) {
				response.UtilResponseReturnJsonFailed(c, pconst.CODE_COMMON_PARAMS_INCOMPLETE)
				return
			}
		}
	}
	param.UserUUID = util.User(c).UUID
	code = service.EditResource(c, param)
	response.UtilResponseReturnJson(c, code, nil)
}

// @Summary DelResource
// @Description 删除ZTA的resource
// @Tags ZTA
// @Produce  json
// @Param id path int true "主键ID"
// @Success 200 {object} controller.Res
// @Router /access/resource/{id} [delete]
func DelResource(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.UtilResponseReturnJsonFailed(c, pconst.CODE_COMMON_PARAMS_INCOMPLETE)
		return
	}
	code := service.DelResource(c, uint64(idInt), util.User(c).UUID)
	response.UtilResponseReturnJson(c, code, nil)
}
