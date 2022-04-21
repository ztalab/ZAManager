package mmysql

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email"`
	UUID  string `json:"uuid" gorm:"column:uuid"`
}

func (User) TableName() string {
	return "zta_user"
}
