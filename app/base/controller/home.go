package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ztalab/ZAManager/pkg/confer"
)

func Welcome(c *gin.Context) {
	now := time.Now().String()
	sysName := confer.ConfigAppGetString("sysname", "default service")
	content := fmt.Sprintf("Welcome to %s@%s", sysName, now)
	c.String(http.StatusOK, content)
}
