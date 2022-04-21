package mparam

import "github.com/ztalab/ZAManager/app/base/mdb"

type GetRelay struct {
	mdb.Paginate
	Name string `json:"name" form:"name"`
}

type AddRelay struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Host    string `json:"host" form:"host" binding:"required"`         // api.github.com
	Port    int    `json:"port" form:"port" binding:"required"`         // 443
	OutPort int    `json:"out_port" form:"out_port" binding:"required"` // 443
}

type EditRelay struct {
	ID      uint64 `json:"id" form:"id" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
	Host    string `json:"host" form:"host" binding:"required"`         // api.github.com
	Port    int    `json:"port" form:"port" binding:"required"`         // 443
	OutPort int    `json:"out_port" form:"out_port" binding:"required"` // 443
}
