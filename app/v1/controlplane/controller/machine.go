package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ztalab/ZAManager/app/v1/controlplane/service"
	"github.com/ztalab/ZAManager/pkg/response"
)

// @Summary LoginUrl
// @Description 根据机器码获取客户端鉴权的url
// @Tags ZTA ControlPlane
// @Produce  json
// @Param machine_id path string true "machine_id"
// @Success 200 {object} controller.Res
// @Router /controlplane/machine/{machine_id} [get]
func LoginUrl(c *gin.Context) {
	code, data := service.GetLoginUrl(c, c.Param("machine_id"))
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary MachineOauth
// @Description 机器鉴权
// @Tags ZTA ControlPlane
// @Produce  json
// @Param hash path string true "hash"
// @Success 200 {object} controller.Res
// @Router /a/{hash} [get]
func MachineOauth(c *gin.Context) {
	service.MachineOauth(c, c.Param("hash"))
}

// @Summary MachineOauth
// @Description 机器鉴权
// @Tags ZTA ControlPlane
// @Produce  json
// @Param category query string true "轮询的主题"
// @Param timeout query int true "超时时间，单位：秒"
// @Success 200 {object} controller.Res
// @Router /machine/auth/poll [get]
func MachineLongpoll(c *gin.Context) {

}
