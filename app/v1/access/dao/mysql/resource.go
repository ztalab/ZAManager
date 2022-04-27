package mysql

import (
	"errors"

	"github.com/ztalab/ZAManager/pkg/util"

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
	if user := util.User(p.c); user != nil {
		query = query.Where("`user_uuid` = ?", user.UUID)
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
	query := orm.Table(p.TableName).Where("id in ?", ids)
	if user := util.User(p.c); user != nil {
		query = query.Where("`user_uuid` = ?", user.UUID)
	}
	err = query.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "GetResourceByIDSli err : %v", err)
	}
	return
}

func (p *Resource) GetResourceByID(id uint64) (info *mmysql.Resource, err error) {
	orm := p.GetOrm()
	query := orm.Table(p.TableName).Where("id = ?", id)
	if user := util.User(p.c); user != nil {
		query = query.Where("`user_uuid` = ?", user.UUID)
	}
	err = query.First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "GetResourceById err : %v", err)
	}
	return
}

func (p *Resource) AddResource(data *mmysql.Resource) (err error) {
	if user := util.User(p.c); user != nil {
		data.UserUUID = user.UUID
	}
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Create(&data).Error
	if err != nil {
		logger.Errorf(p.c, "AddResource err : %v", err)
	}
	return
}

func (p *Resource) EditResource(data *mmysql.Resource) (err error) {
	if user := util.User(p.c); user != nil {
		data.UserUUID = user.UUID
	}
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Save(&data).Error
	if err != nil {
		logger.Errorf(p.c, "EditResource err : %v", err)
	}
	return
}

func (p *Resource) DelResource(id uint64) (err error) {
	orm := p.GetOrm()
	query := orm.Table(p.TableName).Where("id = ?", id)
	if user := util.User(p.c); user != nil {
		query = query.Where("user_uuid = ?", user.UUID)
	}
	err = query.Delete(&mmysql.Resource{}).Error
	if err != nil {
		logger.Errorf(p.c, "DelResource err : %v", err)
	}
	return
}
