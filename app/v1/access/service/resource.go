package service

import (
	"github.com/ztalab/ZAManager/app/v1/access/dao/mysql"
	"github.com/ztalab/ZAManager/app/v1/access/model/mapi"
	"github.com/ztalab/ZAManager/app/v1/access/model/mmysql"
	"github.com/ztalab/ZAManager/app/v1/access/model/mparam"
	"github.com/ztalab/ZAManager/pconst"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func ResourceList(c *gin.Context, param mparam.ResourceList) (code int, ResourceList mapi.ResourceList) {
	count, list, err := mysql.NewResource(c).ResourceList(param)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	ResourceList.List = list
	ResourceList.Paginate.Total = count
	ResourceList.Paginate.PageSize = param.LimitNum
	ResourceList.Paginate.Current = param.Page
	return
}

func AddResource(c *gin.Context, param *mparam.AddResource) (code int, data *mmysql.Resource) {
	data = &mmysql.Resource{
		Name: param.Name,
		UUID: uuid.NewString(),
		Type: param.Type,
		Host: param.Host,
		Port: param.Port,
	}
	err := mysql.NewResource(c).AddResource(data)
	if err != nil {
		return pconst.CODE_COMMON_SERVER_BUSY, nil
	}
	return
}

func EditResource(c *gin.Context, param *mparam.EditResource) (code int) {
	info, err := mysql.NewResource(c).GetResourceByID(param.ID)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if info.ID == 0 {
		code = pconst.CODE_COMMON_DATA_NOT_EXIST
		return
	}
	info.Name = param.Name
	info.Type = param.Type
	info.Host = param.Host
	info.Port = param.Port
	err = mysql.NewResource(c).EditResource(info)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	return
}

func DelResource(c *gin.Context, id uint64) (code int) {
	err := mysql.NewResource(c).DelResource(id)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
	}
	return
}
