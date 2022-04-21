package mmysql

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name string `json:"name"`
	UUID string `json:"uuid" gorm:"column:uuid"`
	Type string `json:"type"`
	Host string `json:"host"`                    // api.github.com
	Port string `json:"port" gorm:"column:port"` // 80-443;3306;6379
}

func (Resource) TableName() string {
	return "zta_resource"
}
