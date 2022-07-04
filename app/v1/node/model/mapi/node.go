package mapi

import (
	"github.com/ztalab/cloudslit/app/base/mapi"
	"github.com/ztalab/cloudslit/app/v1/node/model/mmysql"
)

type NodeList struct {
	List     []mmysql.Node      `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
