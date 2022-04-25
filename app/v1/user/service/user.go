package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ztalab/ZAManager/app/v1/system/dao/mysql"
	"github.com/ztalab/ZAManager/app/v1/system/service"
	"github.com/ztalab/ZAManager/app/v1/user/dao/api"
	userDao "github.com/ztalab/ZAManager/app/v1/user/dao/mysql"
	"github.com/ztalab/ZAManager/app/v1/user/model/mmysql"
	"github.com/ztalab/ZAManager/pconst"
	oauth2Help "github.com/ztalab/ZAManager/pkg/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

func GetRedirectURL(c *gin.Context, company string) (redirectURL string, code int) {
	info, err := mysql.NewOauth2(c).GetOauth2ByCompany(company)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if info.ID == 0 {
		code = pconst.CODE_API_BAD_REQUEST
		return
	}
	config, err := service.Oauth2Config(info)
	if err != nil {
		code = pconst.CODE_API_BAD_REQUEST
		return
	}
	redirectURL, err = oauth2Help.GetOauth2RedirectURL(c, config)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	return
}

func Oauth2Callback(c *gin.Context, company, oauth2Code string) (user *mmysql.User, code int) {
	// 查询对应的配置
	info, err := mysql.NewOauth2(c).GetOauth2ByCompany(company)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if info.ID == 0 {
		code = pconst.CODE_API_BAD_REQUEST
		return
	}
	config, err := service.Oauth2Config(info)
	if err != nil {
		code = pconst.CODE_API_BAD_REQUEST
		return
	}
	switch company {
	case "github":
		githubUser, err := api.GetGithubUser(c, config, oauth2Code)
		if err != nil {
			code = pconst.CODE_API_BAD_REQUEST
			return
		}
		user = &mmysql.User{Email: fmt.Sprintf("%s@github.com", *githubUser.Login), AvatarUrl: *githubUser.AvatarURL}
		if err = userDao.NewUser(c).FirstOrCreateUser(user); err != nil {
			return nil, pconst.CODE_COMMON_SERVER_BUSY
		}
	case "google":
		config.Endpoint = google.Endpoint
	case "facebook":
		config.Endpoint = facebook.Endpoint
	default:
		return nil, pconst.CODE_API_BAD_REQUEST
	}
	return
}
