package mmysql

import (
	"gorm.io/gorm"
)

type Relay struct {
	gorm.Model
	Name    string `json:"name"`
	UUID    string `json:"uuid" gorm:"column:uuid"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	OutPort int    `json:"out_port"`
	CaPem   string `json:"ca_pem"`
	CertPem string `json:"cert_pem"`
}

func (Relay) TableName() string {
	return "zta_relay"
}

type RelayAttrs struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	OutPort int    `json:"out_port"`
	Sort    int    `json:"sort"`
}
