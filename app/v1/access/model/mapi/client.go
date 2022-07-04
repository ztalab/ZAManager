package mapi

import (
	"github.com/ztalab/cloudslit/app/base/mapi"
	"github.com/ztalab/cloudslit/app/v1/access/model/mmysql"
)

type ClientList struct {
	List     []mmysql.Client    `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
