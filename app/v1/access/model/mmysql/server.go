package mmysql

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Server struct {
	gorm.Model
	ResourceID string `json:"resource_id"`
	Name       string `json:"name"`
	UUID       string `json:"uuid" gorm:"column:uuid"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	OutPort    int    `json:"out_port"`
	CaPem      string `json:"ca_pem"`
	CertPem    string `json:"cert_pem"`
}

func (Server) TableName() string {
	return "zta_server"
}

type ServerAttr struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	OutPort int    `json:"out_port"`
}

func (c ServerAttr) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ServerAttr) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
