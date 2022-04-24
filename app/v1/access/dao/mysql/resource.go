package mysql

import (
	"errors"

	"github.com/ztalab/ZAManager/app/v1/access/model/mmysql"
	"github.com/ztalab/ZAManager/app/v1/access/model/mparam"
	"github.com/ztalab/ZAManager/pkg/logger"
	"github.com/ztalab/ZAManager/pkg/mysql"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Resource struct {
	c *gin.Context
	mysql.DaoMysql
}

func NewResource(c *gin.Context) *Resource {
	return &Resource{
		DaoMysql: mysql.DaoMysql{TableName: "zta_resource"},
		c:        c,
	}
}

func (p *Resource) ResourceList(param mparam.ResourceList) (
	total int64, list []mmysql.Resource, err error) {
	orm := p.GetOrm().DB
	query := orm.Table(p.TableName)
	if len(param.Name) > 0 {
		query = query.Where("name like ?", "%"+param.Name+"%")
	}
	if len(param.Type) > 0 {
		query = query.Where("`type` = ?", param.Type)
	}
	if len(param.UserUUID) > 0 {
		query = query.Where("`user_uuid` = ?", param.UserUUID)
	}
	err = query.Model(&list).Count(&total).Error
	if total > 0 {
		offset := param.GetOffset()
		err = query.Limit(param.LimitNum).Offset(offset).
			Order("created_at desc").
			Find(&list).Error
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "ResourceList err : %v", err)
	}
	return
}

func (p *Resource) GetResourceByIDSli(ids []string) (list []mmysql.Resource, err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Where("id in ?", ids).Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "GetResourceByIDSli err : %v", err)
	}
	return
}

func (p *Resource) GetResourceByID(id uint64) (info mmysql.Resource, err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Where("id = ?", id).First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "GetResourceById err : %v", err)
	}
	return
}

func (p *Resource) AddResource(data *mmysql.Resource) (err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Create(&data).Error
	if err != nil {
		logger.Errorf(p.c, "AddResource err : %v", err)
	}
	return
}

func (p *Resource) EditResource(data mmysql.Resource) (err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Save(&data).Error
	if err != nil {
		logger.Errorf(p.c, "EditResource err : %v", err)
	}
	return
}

func (p *Resource) DelResource(id uint64, userUUID string) (err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Where("id = ? and user_uuid = ?", id, userUUID).Delete(&mmysql.Resource{}).Error
	if err != nil {
		logger.Errorf(p.c, "DelResource err : %v", err)
	}
	return
}
