package service

import (
	"strings"
	"time"

	"github.com/ztalab/ZAManager/app/v1/access/dao/api"
	"github.com/ztalab/ZAManager/app/v1/access/dao/mysql"
	"github.com/ztalab/ZAManager/app/v1/access/model/mapi"
	"github.com/ztalab/ZAManager/app/v1/access/model/mmysql"
	"github.com/ztalab/ZAManager/app/v1/access/model/mparam"
	"github.com/ztalab/ZAManager/pconst"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func GetServer(c *gin.Context, param mparam.GetServer) (code int, ServerList mapi.ServerList) {
	count, list, err := mysql.NewServer(c).GetServer(param)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	ServerList.List = list
	ServerList.Paginate.Total = count
	ServerList.Paginate.PageSize = param.LimitNum
	ServerList.Paginate.Current = param.Page
	return
}

func AddServer(c *gin.Context, param *mparam.AddServer) (code int, data *mmysql.Server) {
	data = &mmysql.Server{
		Name:       param.Name,
		ResourceID: param.ResourceID,
		Host:       param.Host,
		Port:       param.Port,
		OutPort:    param.OutPort,
		UUID:       uuid.NewString(),
	}
	attrs := map[string]interface{}{
		"type":     "server",
		"name":     data.Name,
		"host":     data.Host,
		"port":     data.Port,
		"out_port": data.OutPort,
		"uuid":     data.UUID,
	}
	if len(param.ResourceID) > 0 {
		resourceIDSli := strings.Split(strings.TrimSpace(param.ResourceID), ",")
		// 判断传递的资源ID是否合法以及是否存在
		resource, err := mysql.NewResource(c).GetResourceByIDSli(resourceIDSli)
		if err != nil {
			return pconst.CODE_COMMON_DATA_NOT_EXIST, nil
		}
		if len(resourceIDSli) != len(resource) {
			return pconst.CODE_COMMON_DATA_NOT_EXIST, nil
		}
		attrs["resource"] = resource
		sentinelSign, err := api.ApplySign(c, attrs, "zero-access", "zero-access", data.Host, time.Now().AddDate(0, 0, 90))
		if err != nil {
			return pconst.CODE_COMMON_SERVER_BUSY, nil
		}
		data.CaPem = sentinelSign.CaPEM
		data.CertPem = sentinelSign.CertPEM
	}
	err := mysql.NewServer(c).AddServer(data)
	if err != nil {
		return pconst.CODE_COMMON_SERVER_BUSY, nil
	}
	return
}

func EditServer(c *gin.Context, param *mparam.EditServer) (code int) {
	info, err := mysql.NewServer(c).GetServerByID(param.ID)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if info.ID == 0 {
		code = pconst.CODE_COMMON_DATA_NOT_EXIST
		return
	}
	info.Name = param.Name
	info.Host = param.Host
	info.Port = param.Port
	info.OutPort = param.OutPort
	attrs := map[string]interface{}{
		"type":     "server",
		"name":     info.Name,
		"host":     info.Host,
		"port":     info.Port,
		"out_port": info.OutPort,
		"uuid":     info.UUID,
	}
	if len(param.ResourceID) > 0 {
		resourceIDSli := strings.Split(strings.TrimSpace(param.ResourceID), ",")
		// 判断传递的资源ID是否合法以及是否存在
		resource, err := mysql.NewResource(c).GetResourceByIDSli(resourceIDSli)
		if err != nil {
			code = pconst.CODE_COMMON_SERVER_BUSY
			return
		}
		if len(resourceIDSli) != len(resource) {
			code = pconst.CODE_COMMON_SERVER_BUSY
			return
		}
		attrs["resource"] = resource
		sentinelSign, err := api.ApplySign(c, attrs, "zero-access", "zero-access", info.Host, time.Now().AddDate(0, 0, 90))
		if err != nil {
			code = pconst.CODE_COMMON_SERVER_BUSY
			return
		}
		info.CaPem = sentinelSign.CaPEM
		info.CertPem = sentinelSign.CertPEM
	} else {
		info.CaPem = ""
		info.CertPem = ""
	}
	err = mysql.NewServer(c).EditServer(info)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	return
}

func DelServer(c *gin.Context, id uint64) (code int) {
	// TODO 吊销证书
	err := mysql.NewServer(c).DelServer(id)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
	}
	return
}
