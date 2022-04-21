package mapi

import (
	"github.com/ztalab/ZAManager/app/base/mapi"
	"github.com/ztalab/ZAManager/app/v1/access/model/mmysql"
)

type ServerList struct {
	List     []mmysql.Server    `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
