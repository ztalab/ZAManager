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

type Server struct {
	c *gin.Context
	mysql.DaoMysql
}

func NewServer(c *gin.Context) *Server {
	return &Server{
		DaoMysql: mysql.DaoMysql{TableName: "zta_server"},
		c:        c,
	}
}

func (p *Server) ServerList(param mparam.ServerList) (
	total int64, list []mmysql.Server, err error) {
	orm := p.GetOrm().DB
	query := orm.Table(p.TableName)
	if len(param.Name) > 0 {
		query = query.Where("name like ?", "%"+param.Name+"%")
	}
	if param.ResourceID > 0 {
		query = query.Where("find_in_set (?,resource_id)", param.ResourceID)
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
		logger.Errorf(p.c, "ServerList err : %v", err)
	}
	return
}

func (p *Server) GetServerByID(id uint64) (info *mmysql.Server, err error) {
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
		logger.Errorf(p.c, "GetServerById err : %v", err)
	}
	return
}

func (p *Server) AddServer(data *mmysql.Server) (err error) {
	if user := util.User(p.c); user != nil {
		data.UserUUID = user.UUID
	}
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Create(&data).Error
	if err != nil {
		logger.Errorf(p.c, "AddServer err : %v", err)
	}
	return
}

func (p *Server) EditServer(data *mmysql.Server) (err error) {
	if user := util.User(p.c); user != nil {
		data.UserUUID = user.UUID
	}
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Save(&data).Error
	if err != nil {
		logger.Errorf(p.c, "EditServer err : %v", err)
	}
	return
}

func (p *Server) DelServer(id uint64) (err error) {
	orm := p.GetOrm()
	query := orm.Table(p.TableName).Where("id = ?", id)
	if user := util.User(p.c); user != nil {
		query = query.Where("user_uuid = ?", user.UUID)
	}
	err = query.Delete(&mmysql.Server{}).Error
	if err != nil {
		logger.Errorf(p.c, "DelServer err : %v", err)
	}
	return
}
