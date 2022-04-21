package mapi

import (
	"github.com/ztalab/ZAManager/app/base/mapi"
	"github.com/ztalab/ZAManager/app/v1/access/model/mmysql"
)

type ResourceList struct {
	List     []mmysql.Resource  `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
