package mapi

import (
	"github.com/ztalab/ZAManager/app/base/mapi"
	"github.com/ztalab/ZAManager/app/v1/access/model/mmysql"
)

type RelayList struct {
	List     []mmysql.Relay     `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
