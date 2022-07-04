package mapi

import (
	"github.com/ztalab/cloudslit/app/base/mapi"
	"github.com/ztalab/cloudslit/app/v1/access/model/mmysql"
)

type ServerList struct {
	List     []mmysql.Server    `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
