package mapi

import (
	"github.com/ztalab/ZAManager/app/base/mapi"
	"github.com/ztalab/ZAManager/app/v1/node/model/mmysql"
)

type NodeList struct {
	List     []mmysql.Node      `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
