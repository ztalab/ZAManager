package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/ztalab/ZAManager/pconst"

	"github.com/gin-gonic/gin"
	"github.com/ztalab/ZAManager/app/v1/user/service"
)

func Login(c *gin.Context) {
	redirectURL, code := service.GetRedirectURL(c, c.Param("company"))
	if code != pconst.CODE_ERROR_OK {
		// TODO Redirect to one page
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("company %s not support", c.Param("company")))
		return
	}
	c.Redirect(http.StatusSeeOther, redirectURL)
}

func UserDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Hello": "from private", "user": c.GetString("user")})
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
		session.Set("user", user)
		session.Save()
		c.Redirect(http.StatusSeeOther, "/")
	}
	// TODO Redirect to wrong page
}
