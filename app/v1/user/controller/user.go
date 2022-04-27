package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ztalab/ZAManager/pkg/logger"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ztalab/ZAManager/app/v1/user/service"
	"github.com/ztalab/ZAManager/pconst"
	"github.com/ztalab/ZAManager/pkg/response"
	"github.com/ztalab/ZAManager/pkg/util"
)

func Login(c *gin.Context) {
	redirectURL, code := service.GetRedirectURL(c, c.Param("company"))
	if code != pconst.CODE_ERROR_OK {
		// TODO Redirect to BadRequest page
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("company %s not support", c.Param("company")))
		return
	}
	c.Redirect(http.StatusSeeOther, redirectURL)
}

func UserDetail(c *gin.Context) {
	response.UtilResponseReturnJson(c, pconst.CODE_ERROR_OK, util.User(c))
}

func Oauth2Callback(c *gin.Context) {
	session := sessions.Default(c)
	state := session.Get("state")
	if state != c.Query("state") {
		_ = c.AbortWithError(http.StatusUnauthorized, errors.New("state error"))
		return
	}
	if len(c.Query("code")) == 0 {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("code error"))
	}
	user, code := service.Oauth2Callback(c, c.Param("company"), c.Query("code"))
	if code == pconst.CODE_ERROR_OK {
		userBytes, _ := json.Marshal(user)
		session.Set("user", userBytes)
		if err := session.Save(); err != nil {
			logger.Errorf(c, "session save err: %v", err)
			// TODO Redirect to wrong page
			return
		}
		c.Redirect(http.StatusSeeOther, "/")
	}
	// TODO Redirect to wrong page
}
